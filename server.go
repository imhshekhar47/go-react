package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/imhshekhar47/go-react/config"

	"github.com/gin-gonic/gin"
	"github.com/imhshekhar47/go-react/api"
	"github.com/joho/godotenv"
)

var (
	server *gin.Engine
)

func configureRoutes() {
	appModule := server.Group("/api/")
	appModule.GET("/", api.GetAppInfo)

	systemModule := server.Group("/api/system")
	systemModule.GET("/host", api.GetHostInfo)
	systemModule.GET("/pod", api.GetPodInfo)

	healthModule := server.Group("/api/health")
	healthModule.GET("/", api.GetHealth)
}

func main() {
	config := config.GetConfig()

	if err := godotenv.Load(); err == nil {
		log.Println("Custom environment values loaded")
	}

	appHome, exist := os.LookupEnv("APP_HOME")

	if !exist {
		log.Panic("Environment APP_HOME is missing")
	}

	server = gin.Default()

	server.Static("/static", "/app/static")

	server.NoRoute(func(ctx *gin.Context) {
		dir, file := path.Split(ctx.Request.RequestURI)
		ext := filepath.Ext(file)

		if file == "" || ext == "" {
			ctx.File(fmt.Sprintf("%s/%s", appHome, "dist/index.html"))
		} else {
			ctx.File(fmt.Sprintf("%s/%s", appHome, "dist") + path.Join(dir, file))
		}
	})

	configureRoutes()

	if err := server.Run(fmt.Sprintf(":%s", config.Server.Port)); err != nil {
		panic(fmt.Sprintf("Fatal error while launching server: %s\n", err.Error()))
	} else {
		log.Printf("Server listening on port :%s\n", config.Server.Port)
	}
}
