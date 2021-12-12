package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			// Handle error
		}
		c.Set("id", jsonData)
		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.Use(Logger())
	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }

	// receover from panic and write 500 log, high availability
	r.Use(gin.Recovery())

	f, _ := os.Create("gin.log")

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		jsonData, _ := ioutil.ReadAll(param.Request.Body)

		// err = client.Set("id", jsonData, 0).Err()
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s %s \"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
			string(jsonData),
		)
	}))

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

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic YWRtaWx6Q2hhaTNwZWU=' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		log.Println(user)
		// Parse JSON
		var request struct {
			Uid         string `json:"uid" binding:"required"`
			Action      string `json:"action" binding:"required"`
			Category    string `json:"category" binding:"required"`
			SubCategory string `json:"sub_category"`
		}

		if c.Bind(&request) == nil {

			// jsonString, _ := json.Marshal(request)
			// fmt.Println(gin.DefaultWriter, string(jsonString))

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
