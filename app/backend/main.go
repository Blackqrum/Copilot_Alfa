package main

import (
	"alfa-backend/database"
	"alfa-backend/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	db, err := database.Connect("postgres")
	if err != nil {
		log.Fatal("Failed to connect to postgres")
	}
	defer db.Close()

	err = database.CreateDatabase(db)
	if err != nil {
		log.Println("Note: Database might already exist")
	} else {
		log.Println("Database created successfully")
	}

	Appdb, err := database.Connect("myapp")
	if err != nil {
		log.Fatal("Error connecting to myapp database")
	}
	defer Appdb.Close()

	err = database.CreateTables(Appdb)
	if err != nil {
		log.Fatal("Failed to create tables")
	}

	r.POST("/register", handlers.RegisterHandler(Appdb))
	r.POST("/api/assistant", handlers.BusinessAssistantHandler)

	r.Run(":8080")
}
