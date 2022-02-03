package models

type EquipmentResponse struct {
	Id			 string `json:"id"`
	Name         string  `json:"name"`
	Type         string `json:"type"`
	SerialNumber string `json:"serialNumber"`
}
