package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Question struct {
	ID      int      `json:"id"`
	Text    string   `json:"text"`
	Choices []string `json:"choices"`
	Answer  int      `json:"-"` // Index of the correct answer (hidden from response)
}

type Response struct {
	QuestionID int `json:"question_id"`
	Answer     int `json:"answer"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var users = []User{}
var userResponses = make(map[string]map[int]int) // Store latest answer for each question
var jwtKey = []byte("secret_key")

var questions = []Question{
	{ID: 1, Text: "What is the capital of France?", Choices: []string{"Paris", "London", "Berlin", "Madrid"}, Answer: 0},
	{ID: 2, Text: "What is 2 + 2?", Choices: []string{"3", "4", "5", "6"}, Answer: 1},
	{ID: 3, Text: "Which planet is known as the Red Planet?", Choices: []string{"Earth", "Mars", "Jupiter", "Venus"}, Answer: 1},
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
			claims := &Claims{Email: loginReq.Email, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}
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
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) { return jwtKey, nil })
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Next()
	}
}

func getQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, questions)
}

func submitResponse(c *gin.Context) {
	email, _ := c.Get("email")
	var response Response
	if err := c.ShouldBindJSON(&response); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userResponses[email.(string)] == nil {
		userResponses[email.(string)] = make(map[int]int)
	}
	userResponses[email.(string)][response.QuestionID] = response.Answer // Override previous answer
	c.JSON(http.StatusOK, gin.H{"message": "Response recorded"})
}

func getResult(c *gin.Context) {
	email, _ := c.Get("email")
	responses, exists := userResponses[email.(string)]
	if !exists {
		c.JSON(http.StatusOK, gin.H{"score": 0, "message": "No responses found"})
		return
	}
	result := make(map[int]bool)
	for _, q := range questions {
		result[q.ID] = false
		if answer, ok := responses[q.ID]; ok && answer == q.Answer {
			result[q.ID] = true
		}
	}
	c.JSON(http.StatusOK, gin.H{"score": result})
}

func main() {
	r := gin.Default()

	r.Use(corsMiddleware())

	r.POST("/register", registerUser)
	r.POST("/login", loginUser)

	protected := r.Group("/").Use(authenticate())
	protected.GET("/questions", getQuestions)
	protected.POST("/response", submitResponse)
	protected.GET("/result", getResult)

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
