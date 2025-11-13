package main

import (
	"fmt"
	//"database/sql"
	"alfa-backend/database"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	//"log"
	"net/http"
)

func main() {
	r := gin.Default()

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect")
	}

	err = database.CreateDatabase(db)
	if err != nil {
		log.Fatal("Failed to launch")
	}

	defer db.Close()

	fmt.Println("Database created successfully")

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello nigga")
	})

	r.Run(":8080")
}
