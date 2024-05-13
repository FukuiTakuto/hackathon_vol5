package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supabase-community/supabase-go"
)

func main() {
    supabaseURL := 
    supabaseKey := 
    sb,err := supabase.NewClient(supabaseURL, supabaseKey,nil)
		if err != nil {
			panic(err) 
	}

    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        // データベースからランダムな1行のユーザ情報を取得
				query := sb.From("Users").Select("username", "email",false).Order("RANDOM()").Limit(1,"")

				res, err := query.Get(c)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        var user struct {
            Username string `json:"username"`
            Email    string `json:"email"`
        }
        if err := res.Scan(&user); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.HTML(http.StatusOK, "index.html", gin.H{
            "username": user.Username,
            "email":    user.Email,
        })
    })

    r.Run(":8080")
}
