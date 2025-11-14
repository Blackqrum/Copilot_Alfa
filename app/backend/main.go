package main

import (
	"fmt"
	//"database/sql"
	"alfa-backend/database"
	"alfa-backend/handlers"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	//"log"
)

func main() {
	r := gin.Default()

	db, err := database.Connect("postgres")
	if err != nil {
		log.Fatal("Failed to connect")
	}

	defer db.Close()

	err = database.CreateDatabase(db)
	if err != nil {
		log.Println("Note: Database might already exist")
	} else {
		fmt.Println("Database created successfully")
	}

	fmt.Println("Database created successfully")

	Appdb, err := database.Connect("myapp")

	if err != nil {
		log.Fatal("Error with connection")
	}

	defer Appdb.Close()

	err = database.CreateTables(Appdb)

	if err != nil {
		log.Fatal("Failed to create")
	}

	fmt.Println("✅ Server ready!") // ← ЭТО ДОЛЖНО ВЫВЕСТИСЯ

	r.POST("/register", handlers.RegisterHandler(Appdb))
	r.POST("/login", handlers.LoginHandler(Appdb))

	r.Run(":8080")
}
