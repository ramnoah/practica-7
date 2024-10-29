package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string
	Email string
}

func main() {
	r := gin.Default()

	user := &User{
		Name:  "Noe",
		Email: "ramnoah@gmail.com",
	}

	users := []User{
		{
			Name:  "Mario",
			Email: "mar@mail.com",
		},
		{
			Name:  "Fer",
			Email: "fer@gmail.com",
		},
	}

	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":      "Main website",
			"totalUsers": 11,
			"user":       user,
			"users":      users,
		})
	})
	r.Run(":8001")
	fmt.Println("Running")
}
