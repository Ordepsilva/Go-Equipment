package models

import "equipment/internal/equipment/domain"

type Equipment struct {
	Name         string  `json:"name"`
	Type         string `json:"type"`
	SerialNumber string `json:"serialNumber"`
}

func  (equipment Equipment) ToDomain (ID int) domain.Equipment{
	return domain.Equipment{
		ID:           ID,
		Name:         equipment.Name,
		Type:         equipment.Type,
		SerialNumber: equipment.SerialNumber,
	}
}