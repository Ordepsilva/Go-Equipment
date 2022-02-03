package api

import (
	"equipment/api/models"
	"equipment/internal/equipment"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"net/http"
	"strconv"
)

type EquipmentAPI struct {
	equipmentService equipment.Service
}

func NewEquipmentAPI(neo4jSession neo4j.Session) *EquipmentAPI {
	return &EquipmentAPI{equipmentService: equipment.NewService(neo4jSession)}
}

func (equipmentAPI EquipmentAPI) InitEquipmentAPI(router *gin.RouterGroup) {
	routerGroup := router.Group("/equipment")
	routerGroup.POST("", equipmentAPI.AddEquipment)
	routerGroup.GET("", equipmentAPI.GetEquipments)
	routerGroup.GET("/:id", equipmentAPI.GetEquipmentByID)
	routerGroup.DELETE("/:id", equipmentAPI.DeleteEquipment)
	routerGroup.PUT("/:id", equipmentAPI.UpdateEquipment)

}

func (equipmentAPI EquipmentAPI) AddEquipment(context *gin.Context) {
	var equipment models.Equipment
	err := context.ShouldBind(&equipment)

	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	createdEquipmentID, err := equipmentAPI.equipmentService.AddEquipment(equipment.Name, equipment.Type, equipment.SerialNumber)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusCreated, createdEquipmentID)
	return
}

func (equipmentAPI EquipmentAPI) GetEquipments(context *gin.Context) {
	equipments, err := equipmentAPI.equipmentService.GetEquipments()
	if err != nil {
		context.Status(http.StatusInternalServerError)
	}

	context.JSON(http.StatusOK, equipments)
}

func (equipmentAPI EquipmentAPI) GetEquipmentByID(context *gin.Context) {
	equipmentID := context.Param("id")
	id, err := strconv.Atoi(equipmentID)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	equipment, err := equipmentAPI.equipmentService.GetEquipmentByID(id)
	if err != nil {
		context.Status(http.StatusNotFound)
		return
	}

	context.JSON(http.StatusOK, equipment)
}

func (equipmentAPI EquipmentAPI) DeleteEquipment(ctx *gin.Context) {
	equipmentID := ctx.Param("id")
	id, err := strconv.Atoi(equipmentID)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = equipmentAPI.equipmentService.DeleteEquipment(id)
	if err == nil {
		ctx.Status(http.StatusOK)
		return
	}
	ctx.Status(http.StatusNotFound)

}

func (equipmentAPI EquipmentAPI) UpdateEquipment(ctx *gin.Context) {
	equipmentID := ctx.Param("id")
	id, err := strconv.Atoi(equipmentID)

	var equipment models.Equipment
	if err := ctx.ShouldBind(&equipment); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	equipmentDomain := equipment.ToDomain(id)

	err = equipmentAPI.equipmentService.UpdateEquipment(&equipmentDomain)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.Status(http.StatusOK)
}
