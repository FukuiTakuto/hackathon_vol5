package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
)

func main() {
	router := gin.Default()

	router.GET("/update/:id", updateUser)

	router.Run(":8080")
}

type User struct {
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
}

func updateUser(c *gin.Context) {
	supabaseUrl := ""
	supabaseKey := ""
	client := supa.CreateClient(supabaseUrl, supabaseKey)

	userID := c.Query("id")
	if userID == "" {
		  c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
	}

	updatedData := User{
			Name:  "New Name",  
			Hobby: "New Hobby", 
	}

	var results map[string]interface{}
	err := client.DB.From("users").Update(updatedData).Eq("id", userID).Execute(&results)
	if err != nil {
		  c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "results": results})
}
