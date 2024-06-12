package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddSupermarket(c *gin.Context) {
	name := c.Param("name")
	street := c.Param("street")
	postalCode := c.Param("postal_code")
	city := c.Param("city")
	houseNumber := c.Param("houseNumber")

	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert new supermarket entry
	query := `INSERT INTO Supermarkets (name, street, postal_code, city, house_number) 
	          VALUES (?, ?, ?, ?, ?)`
	_, err = db.Exec(query, name, street, postalCode, city, houseNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add supermarket: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supermarket added successfully!"})
}
