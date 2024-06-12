package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Wine represents the wine data structure
type Wine struct {
	Name          string          `json:"name"`
	Year          int             `json:"year"`
	Volume        float64         `json:"volume"`
	CheapestPrice sql.NullFloat64 `json:"cheapestPrice"`
}

// Database connection function
func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// getWine handles GET requests for retrieving wine data by ID
func GetWine(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wine ID"})
		return
	}

	db, err := connectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	var wine Wine
	query := `
		SELECT Wine.name, Wine.year, Wine.volume, MIN(WineSupermarkets.price) as cheapestPrice
		FROM Wine
		LEFT JOIN WineSupermarkets ON Wine.id = WineSupermarkets.wine_id
		WHERE Wine.id = ?
		GROUP BY Wine.id, Wine.name, Wine.year, Wine.volume`
	row := db.QueryRow(query, id)
	err = row.Scan(&wine.Name, &wine.Year, &wine.Volume, &wine.CheapestPrice)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wine not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to query the database: %v", err)})
		return
	}

	response := gin.H{
		"name":   wine.Name,
		"year":   wine.Year,
		"volume": wine.Volume,
	}

	if wine.CheapestPrice.Valid {
		response["cheapestPrice"] = wine.CheapestPrice.Float64
	} else {
		response["cheapestPrice"] = nil
	}

	c.JSON(http.StatusOK, response)
}
