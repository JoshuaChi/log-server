package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

var client *redis.Client

func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	gin.DisableConsoleColor()
	r := gin.Default()

	// receover from panic and write 500 log, high availability
	r.Use(gin.Recovery())

	f, _ := os.Create("logs/gin.log")

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.POST("/log", TokenAuthMiddleware(), Log)
	r.POST("/login", Login)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	log.Fatal(r.Run(":8080"))
}
