package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddSupermarket(c *gin.Context) {
	name := c.Query("name")
	street := c.Query("street")
	postalCode := c.Query("postal_code")
	city := c.Query("city")
	houseNumber := c.Query("houseNumber")

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

func setWineSupermarket(c *gin.Context) {
	wineID, err := strconv.Atoi(c.Query("wine_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wineID"})
		return
	}
	supermarketID, err := strconv.Atoi(c.Query("supermarket_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid supermarketID"})
		return
	}

	priceStr := c.Query("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
		return
	}

	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	// Print values for debugging
	fmt.Printf("WineID: %d, SupermarketID: %d, Price: %f\n", wineID, supermarketID, price)

	// Insert or update the price of the wine in the supermarket
	query := `
	INSERT INTO WineSupermarkets (wine_id, supermarket_id, price)
	VALUES (?, ?, ?)
	ON CONFLICT(wine_id, supermarket_id) DO UPDATE SET price=excluded.price`
	_, err = db.Exec(query, wineID, supermarketID, price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to execute query: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wine price set successfully"})
}
