package domain

import (
	"errors"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
)

type Equipment struct {
	ID           int
	Name         string `json:"name"`
	Type         string
	SerialNumber string
}

func NewEquipment(name string, equipmentType string, serialNumber string) (*Equipment, error) {

	if name == "" {
		return nil, errors.New("name is empty")
	}

	if equipmentType == "" {
		return nil, errors.New("equipmentType is empty")
	}

	if serialNumber == "" {
		return nil, errors.New("serialNumber is empty")
	}

	return &Equipment{Name: name, Type: equipmentType, SerialNumber: serialNumber}, nil
}

func (equip *Equipment) Update(updatedEquipment *Equipment) error {

	if updatedEquipment.Name == "" {
		return errors.New("name is empty")
	}

	if updatedEquipment.Type == "" {
		return errors.New("equipmentType is empty")
	}

	if updatedEquipment.SerialNumber == "" {
		return errors.New("serialNumber is empty")
	}

	equip.Name = updatedEquipment.Name
	equip.Type = updatedEquipment.Type
	equip.SerialNumber = updatedEquipment.SerialNumber

	return nil
}

func FromRepo(record *db.Record) Equipment {
	id, _ := record.Get("id")
	name, _ := record.Get("name")
	equipmentType, _ := record.Get("type")
	serialNumber, _ := record.Get("serialnumber")

	return Equipment{
		ID:           int(id.(int64)),
		Name:         name.(string),
		Type:         equipmentType.(string),
		SerialNumber: serialNumber.(string),
	}
}
