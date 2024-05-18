package main

import (
	"fmt"
	"net/http"

	supa "github.com/nedpals/supabase-go"
)

type User struct {
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	supabaseUrl := "<SUPABASE-URL>"
	supabaseKey := "<SUPABASE-KEY>"
	client := supa.CreateClient(supabaseUrl, supabaseKey)

	userID := r.URL.Query().Get("id")
	if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
	}

	updatedData := User{
			Name:  "New Name",  
			Hobby: "New Hobby", 
	}

	var results map[string]interface{}
	err := client.DB.From("users").Update(updatedData).Eq("id", userID).Execute(&results)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	fmt.Fprintf(w, "User updated successfully: %v", results)
}
