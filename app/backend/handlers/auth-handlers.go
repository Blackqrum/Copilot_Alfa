package handlers

import (
	"alfa-backend/auth"
	"alfa-backend/database"
	"alfa-backend/modules"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func LoginHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req modules.LoginRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		fmt.Println("Login attempt for email:", req.Email)

		user, err := database.GetEmailFromDb(db, req.Email)
		fmt.Println("GetEmailFromDb result - user:", user, "error:", err)

		if err != nil {
			fmt.Println("User not found or DB error")
			c.JSON(401, gin.H{"error": "Invalid credentials"})
			return
		}

		fmt.Println("Checking password...")
		if !auth.CheckPasswordHash(req.Password, user.Password) {
			c.JSON(401, gin.H{"error": "Invalid credentials"})
			return
		}

		c.JSON(200, gin.H{"message": "Successful entry"})
	}
}

func RegisterHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req modules.RegisterRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		hashed_password, err := auth.HashPassword(req.Password)
		if err != nil {
			c.JSON(500, "error processing password")
			return
		}

		err = database.AddUserToDB(db, req.Email, hashed_password, req.Name)
		if err != nil {
			log.Println("AddUserToDB error:", err)
			if strings.Contains(err.Error(), "unique constraint") || strings.Contains(err.Error(), "повторяющееся значение ключа") {
				c.JSON(400, "user already exists")
				return
			}
			c.JSON(500, gin.H{"error": "problem with database"})
			return
		}

		c.JSON(200, gin.H{"message": "user added successfully"})

	}
}
