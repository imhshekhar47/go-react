package api

import (
	"net/http"

	"github.com/imhshekhar47/go-react/config"

	"github.com/gin-gonic/gin"
)

// AppInfo : data model
type AppInfo struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Developer string `json:"developer"`
}

// GetAppInfo : returns application information
func GetAppInfo(ctx *gin.Context) {
	// appName, exists := os.LookupEnv("APP_NAME")
	// if !exists {
	// 	appName = "go-react"
	// }
	// appVersion, exists := os.LookupEnv("APP_VERSION")
	// if !exists {
	// 	appVersion = "1.0.0"
	// }

	// appDeveloper, exists := os.LookupEnv("APP_DEVELOPER")
	// if !exists {
	// 	appDeveloper = "imhshekhar47"
	// }

	config := config.GetConfig()
	ctx.JSON(http.StatusOK, AppInfo{
		Name:      config.Application.Name,
		Version:   config.Application.Version,
		Developer: config.Application.Developer,
	})
}
