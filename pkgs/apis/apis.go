package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/joshuachi/logserver/pkgs/auth"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var user User

func init() {
	err := godotenv.Load() // The Original basic .env which is shared cross ENVs
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user = User{
		ID:       1,
		Username: os.Getenv("DEVELOPER_NAME"),
		Password: os.Getenv("DEVELOPER_PASSWORD"),
	}
}

func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	ts, err := auth.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	saveErr := auth.CreateAuth(user.ID, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}

func Log(c *gin.Context) {
	var request struct {
		Uid         string `json:"uid" binding:"required"`
		Action      string `json:"action" binding:"required"`
		Category    string `json:"category" binding:"required"`
		SubCategory string `json:"sub_category"`
	}

	if c.Bind(&request) == nil {
		jsonString, _ := json.Marshal(request)
		_log(c, string(jsonString))
		c.JSON(http.StatusOK, gin.H{})
	}
}

func _log(c *gin.Context, msg string) {
	start := time.Now()
	path := c.Request.URL.Path
	// raw := c.Request.URL.RawQuery
	clientIP := c.ClientIP()
	statusCode := c.Writer.Status()
	bodySize := c.Writer.Size()

	r := fmt.Sprintf("[%s] - [\"%s %s\" - %d - %d] - %s\n", start.Format(time.RFC3339), clientIP, path, statusCode, bodySize, msg)
	gin.DefaultWriter.Write([]byte(r))

}
