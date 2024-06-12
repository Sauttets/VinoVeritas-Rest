package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func AddWine(c *gin.Context) {
	name := c.Query("name")
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year"})
		return
	}
	country := c.Query("country")
	wineType := c.Query("type")
	description := c.Query("description")
	imageURL := c.Query("imageURL")
	volume, err := strconv.ParseFloat(c.Query("volume"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid volume"})
		return
	}
	volAlc, err := strconv.ParseFloat(c.Query("volAlc"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid volAlc"})
		return
	}

	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert new wine entry
	query := `INSERT INTO Wine (name, year, country, type, description, imageURL, volume, volAlc) 
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, name, year, country, wineType, description, imageURL, volume, volAlc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add wine: %v", err)})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve last insert ID: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wine added successfully!", "id": id})
}
