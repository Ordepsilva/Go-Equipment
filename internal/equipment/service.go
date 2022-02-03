package equipment

import (
	"equipment/internal/equipment/domain"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Service struct {
	repository Repository
}

func NewService(neo4jSession neo4j.Session) Service {
	return Service{
		repository: NewRepository(neo4jSession),
	}
}

func (service Service) AddEquipment(name string, equipmentType string, serialNumber string) (string, error) {
	newEquipment, err := domain.NewEquipment(name, equipmentType, serialNumber)
	if err != nil {
		return "", err
	}

	recordID, err := service.repository.CreateEquipment(*newEquipment)
	if err != nil {
		return "", err
	}

	return recordID, nil
}

func (service Service) GetEquipments() (*[]domain.Equipment, error) {
	equipments, err := service.repository.GetEquipments()
	if err != nil {
		return nil, err
	}

	return equipments, nil
}

func (service Service) GetEquipmentByID(equipmentID int) (*domain.Equipment, error) {

	return service.repository.GetEquipmentByID(equipmentID)
}

func (service Service) DeleteEquipment(equipmentID int) error {
	_, err := service.GetEquipmentByID(equipmentID)
	if err != nil {
		return err
	}
	return service.repository.DeleteEquipment(equipmentID)
}

func (service Service) UpdateEquipment(updatedEquipment *domain.Equipment) error {

	equipment, err := service.GetEquipmentByID(updatedEquipment.ID)
	if err != nil {
		return err
	}
	if equipment == nil{
		return errors.New("Equipment not found!")
	}

	err = service.repository.UpdateEquipment(*updatedEquipment)
	return err
}
