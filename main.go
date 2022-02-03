package main

import (
	"equipment/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	router := gin.Default()
	bdUrl := "neo4j://localhost:7687"
	username := "test"
	password := "test"

	driver, err := neo4j.NewDriver(bdUrl, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	equipmentAPI := api.NewEquipmentAPI(session)
	equipmentAPI.InitEquipmentAPI(router.Group(
		""))
	err = router.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
