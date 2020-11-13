package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health : model
type Health struct {
	Version string `json:"version"`
	Status  string `json:"status"`
}

// GetHealth : Returns health status of the api
func GetHealth(ctx *gin.Context) {
	data := Health{
		Version: "1.0.0",
		Status:  "OK",
	}
	ctx.JSON(http.StatusOK, data)
}
