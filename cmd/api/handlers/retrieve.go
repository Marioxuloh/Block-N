package handlers

import (
	"Block-N/cmd/gRPC"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RetrieveParams struct {
	Domain string `json:"domain"`
}

func (s Handler) Retrieve(c *gin.Context) {

	// obtenemos y validamos mediante json el tipo de dato
	var params RetrieveParams
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// llamamos al servicio create a traves del puerto
	neighbor, err := gRPC.Retrieve(s.Node, params.Domain, 0)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
	}

	c.JSON(http.StatusAccepted, gin.H{"id": neighbor.Id, "address": neighbor.Address})

}
