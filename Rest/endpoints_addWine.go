package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddWine(c *gin.Context) {
	name := c.Param("name")
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year"})
		return
	}
	country := c.Param("country")
	wineType := c.Param("type")
	description := c.Param("description")
	imageURL := c.Param("imageURL")
	volume, err := strconv.ParseFloat(c.Param("volume"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid volume"})
		return
	}
	volAlc, err := strconv.ParseFloat(c.Param("volAlc"), 64)
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
	_, err = db.Exec(query, name, year, country, wineType, description, imageURL, volume, volAlc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add wine: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wine added successfully!"})
}
