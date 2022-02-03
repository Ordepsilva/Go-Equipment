package equipment

import (
	"equipment/internal/equipment/domain"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"strconv"
)

type Repository struct {
	Neo4jSession neo4j.Session
}

func NewRepository(neo4jSession neo4j.Session) Repository {
	return Repository{
		Neo4jSession: neo4jSession,
	}
}

func (repo *Repository) CreateEquipment(equipment domain.Equipment) (string, error) {
	var recordID string
	query := "CREATE (n:Equipment {name: $name, serialnumber: $serialnumber, type: $type}) RETURN id(n) as id"
	params := map[string]interface{}{
		"name":         equipment.Name,
		"serialnumber": equipment.SerialNumber,
		"type":         equipment.Type,
	}
	result, err := repo.Neo4jSession.Run(query, params)

	if err != nil {
		return "", err
	}

	for result.Next() {
		id, _ := result.Record().Get("id")
		recordID = strconv.Itoa(int(id.(int64)))
	}

	return recordID, nil
}
func (repo *Repository) GetEquipments() (*[]domain.Equipment, error) {
	objects := &[]domain.Equipment{}
	query := "Match (n:Equipment) return n.name as name, n.serialnumber as serialnumber, n.type as type, ID(n) as id "

	result, err := repo.Neo4jSession.Run(query, nil)

	if err != nil {
		return nil, err
	}

	for result.Next() {
		*objects = append(*objects, domain.FromRepo(result.Record()))
	}

	return objects, nil
}

func (repo *Repository) GetEquipmentByID(equipmentID int) (*domain.Equipment, error) {
	var equipment domain.Equipment
	query := "Match (n:Equipment) Where ID(n)= $recordID  Return n.name as name, n.serialnumber as serialnumber, n.type as type, ID(n) as id "
	params := map[string]interface{}{
		"recordID": equipmentID,
	}

	result, err := repo.Neo4jSession.Run(query, params)
	if err != nil {
		return nil, result.Err()
	}

	if result.Next() {
		equipment = domain.FromRepo(result.Record())
		return &equipment, nil
	}

	return nil, errors.New("not found")
}

func (repo *Repository) DeleteEquipment(equipmentID int) error {
	query := "Match (n:Equipment) Where ID(n)= $recordID  Delete n"
	params := map[string]interface{}{
		"recordID": equipmentID,
	}
	result, err := repo.Neo4jSession.Run(query, params)
	if err != nil {
		return err
	}

	summary, err := result.Consume()
	if err != nil {
		return err
	}

	counters := summary.Counters()
	count := counters.NodesDeleted()
	if count > 0 {
		return nil
	}

	return errors.New("Failed to delete item")
}

func (repo *Repository) UpdateEquipment(updatedEquip domain.Equipment) error {
	query := "Match (n:Equipment) Where ID(n)= $recordID  SET n = {name: $name, serialnumber: $serialnumber, type: $type} Return ID(n) as id"
	params := map[string]interface{}{
		"recordID":     updatedEquip.ID,
		"name":         updatedEquip.Name,
		"serialnumber": updatedEquip.SerialNumber,
		"type":         updatedEquip.Type,
	}

	result, err := repo.Neo4jSession.Run(query, params)
	if err != nil {
		return err
	}

	summary, err := result.Consume()
	if err != nil {
		return err
	}

	counters := summary.Counters()

	if counters.PropertiesSet() > 0{
		return nil
	}

	return errors.New("Failed to update equipment")
}
