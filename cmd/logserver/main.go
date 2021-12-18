package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/joshuachi/logserver/pkgs/apis"
	"github.com/joshuachi/logserver/pkgs/auth"
)

func init() {

	err := godotenv.Load() // The Original basic .env which is shared cross ENVs
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("ENV")

	if len(env) < 1 {
		env = "dev"
	}

	godotenv.Load(".env." + env)
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	gin.DisableConsoleColor()
	r := gin.Default()

	// receover from panic and write 500 log, high availability
	r.Use(gin.Recovery())

	f, _ := os.Create("logs/" + os.Getenv("LOG_FILENAME"))

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.POST("/api/v1/log", auth.TokenAuthMiddleware(), apis.Log)
	r.POST("/api/v1/login", apis.Login)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}
