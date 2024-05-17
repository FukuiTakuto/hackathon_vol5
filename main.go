package main

import (
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/gin-gonic/gin"
)

type User struct {
    UserID string `json:"userId"`
    Hobby  string `json:"hobby"`
    Age    int    `json:"age"`
}

func main(){
	router:=gin.Default()

	router.PUT("/update:id",UpdateProfileHandler)

    router.Run(":8080")
}

func UpdateProfileHandler(c*gin.Context){
    userID := c.Param("id")

    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    supabaseURL := ""
    apiKey := ""

    updateUser := map[string]interface{}{
        "hobby": user.Hobby,
        "age":   user.Age,
    }

    client := resty.New()
    resp, err := client.R().
        SetHeader("apikey", apiKey).
        SetHeader("Authorization", "Bearer "+apiKey).
        SetHeader("Content-Type", "application/json").
        SetBody(updateUser).
        Patch(supabaseURL + "/rest/v1/User?id=eq." + userID)

    if err != nil {
        log.Println("Error updating user:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    if resp.StatusCode() != http.StatusOK {
        log.Println("Supabase responded with status:", resp.Status())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}