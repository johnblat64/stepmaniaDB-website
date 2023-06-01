package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// get environment variable for backend api endpoint and crash if not set with log
	GEnvironmentConfig.BackendApiUrl = os.Getenv("BACKEND_API_ENDPOINT")
	if GEnvironmentConfig.BackendApiUrl == "" {
		log.Fatal("BACKEND_API_ENDPOINT environment variable not set")
	}

	GEnvironmentConfig.FileStoreUrl = os.Getenv("FILE_STORE_ENDPOINT")
	if GEnvironmentConfig.FileStoreUrl == "" {
		log.Fatal("FILE_STORE_ENDPOINT environment variable not set")
	}

	r := gin.Default()
	r.Static("/resources", "./resources")
	r.GET("/", getSongs)
	r.GET("/songs", getSongs)
	r.GET("/songs/:songid", getSong)
	r.GET("/packs/:packid", getPack)
	r.GET("/about", getAbout)
	r.GET("/todo", getTodo)
	r.NoRoute(getNotFound)

	r.Run(":80")
}
