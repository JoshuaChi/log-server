package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// receover from panic and write 500 log, high availability
	r.Use(gin.Recovery())

	f, _ := os.Create("gin.log")

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "Chai3pee", // user:foo password:bar
		"dev1":  "cooLiet6", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   echo -n 'admin:Chai3pee'  | openssl base64  // YWRtaWx6Q2hhaTNwZWU=

		curl -i -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic YWRtaW46Q2hhaTNwZWU=' \
	  	-H 'content-type: application/json' \
	  	-d '{"Uid":"1E43", "Action":"Shop", "Category":"product", "SubCategory":"test"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		// user := c.MustGet(gin.AuthUserKey).(string)

		// log.Println(user)
		// Parse JSON
		var request struct {
			Uid         string `json:"uid" binding:"required"`
			Action      string `json:"action" binding:"required"`
			Category    string `json:"category" binding:"required"`
			SubCategory string `json:"sub_category"`
		}

		if c.Bind(&request) == nil {

			jsonString, _ := json.Marshal(request)
			log(c, string(jsonString))
			c.JSON(http.StatusOK, gin.H{})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func log(c *gin.Context, msg string) {
	start := time.Now()
	path := c.Request.URL.Path
	// raw := c.Request.URL.RawQuery
	clientIP := c.ClientIP()
	statusCode := c.Writer.Status()
	bodySize := c.Writer.Size()

	r := fmt.Sprintf("[%s] - [\"%s %s\" - %d - %d] - %s\n", start.Format(time.RFC3339), clientIP, path, statusCode, bodySize, msg)
	gin.DefaultWriter.Write([]byte(r))

}
