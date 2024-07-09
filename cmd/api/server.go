package api

import (
	"Block-N/cmd/api/handlers"
	"Block-N/services/node"
	"log"

	"github.com/gin-gonic/gin"
)

func InitRESTServer(n *node.Node) error {

	gin.SetMode(gin.ReleaseMode)

	ginEngine := gin.New()

	// habilitar capa de proteccion para errores inesperados
	ginEngine.Use(gin.Recovery())

	// habilitar el logger para las llamadas
	ginEngine.Use(gin.Logger())

	retrieveHandler := handlers.Handler{
		Node: n,
	}

	ginEngine.GET("/retrieve", retrieveHandler.Retrieve)

	log.Printf("server REST listening at 192.168.1.144:8080")

	err := ginEngine.Run("192.168.1.144:8080")
	if err != nil {
		return err
	}

	return nil
}
