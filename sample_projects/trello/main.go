package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Task struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Stage string `json:"stage"`
	Order int    `json:"order"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var tasks = []Task{
	{ID: 1, Text: "Sample Task", Stage: "Todo", Order: 1},
}
var users = []User{}
var nextID = 2
var stages = []string{"Todo", "On hold", "Inprogress", "Done"}
var jwtKey = []byte("secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func createTask(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask.ID = nextID
	nextID++
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func moveTask(c *gin.Context) {
	var req struct {
		ID    int    `json:"id"`
		Stage string `json:"stage"`
		Order int    `json:"order"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == req.ID {
			tasks[i].Stage = req.Stage
			tasks[i].Order = req.Order
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func registerUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users = append(users, newUser)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func loginUser(c *gin.Context) {
	var loginReq User
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, user := range users {
		if user.Email == loginReq.Email && user.Password == loginReq.Password {
			expirationTime := time.Now().Add(24 * time.Hour)
			claims := &Claims{
				Email: loginReq.Email,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, _ := token.SignedString(jwtKey)

			c.JSON(http.StatusOK, gin.H{"token": tokenString})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

func authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}
		tokenString := strings.Split(tokenStr, " ")[1]

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(corsMiddleware())

	r.POST("/register", registerUser)
	r.POST("/login", loginUser)

	protected := r.Group("/").Use(authenticate())
	protected.POST("/tasks", createTask)
	protected.GET("/tasks", getTasks)
	protected.PUT("/tasks/move", moveTask)

	r.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
