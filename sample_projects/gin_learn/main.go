package main

import (
	"gin_learn/common"
	"gin_learn/student"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	r := gin.Default()
	// Add middleware for all the routes
	// r.Use(common.Logger())

	envFile, _ := godotenv.Read(".env")

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	v1 := r.Group("/api")
	student.StudentsRegister(v1.Group("/students"))

	// Add a middleware to a specific route
	r.GET("/ping", common.Logger(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":" + envFile["PORT"])
}
