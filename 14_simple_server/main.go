package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// User represents a user
// JSON output example:
//
//	{
//	 "id":"123",
//	 "name":"John Doe",
//	 "email":"john@mail.com"
//	}
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name" binding:"required,min=1,max=100"`
	Email string `json:"email" binding:"required,email"`
}

var (
	// Simulated an in-memory database
	usersTable      = make(map[string]User)
	userIDIncrement = 0
)

func newUserID() string {
	userIDIncrement += 1
	return strconv.Itoa(userIDIncrement)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	// should validate id

	user, exists := usersTable[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = newUserID()
	usersTable[user.ID] = user

	c.JSON(http.StatusCreated, user)
}

func updateUser(c *gin.Context) {
	// your code here
	// ....
}

func listUsers(c *gin.Context) {
	// your code here
	// ....
}

func main() {
	r := gin.Default()

	r.GET("/users/:id", getUser)
	r.POST("/users", createUser)

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func demoJSONUnmarshal() {
	jsonStr := `{"id":"1","name":"Alice","email":"alice@mail.com"}`

	var user User
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		log.Printf("Failed to unmarshal JSON: %v", err)
		return
	}
	log.Printf("Unmarshaled User: %+v\n", user)
}

func demoJsonMarshal() {
	user := User{
		ID:    "2",
		Name:  "Bob",
		Email: "bob@mail.com",
	}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		log.Printf("Failed to marshal User to JSON: %v", err)
		return
	}
	log.Printf("Marshaled JSON: %s\n", string(jsonBytes))
}
