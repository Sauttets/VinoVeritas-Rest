package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddFitsTo(c *gin.Context) {
	name := c.Query("fit_id")

	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert new fitsTo entry
	query := `INSERT INTO FitsTo (description) VALUES (?)`
	result, err := db.Exec(query, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add fitsTo: %v", err)})
		return
	}

	// Get the ID of the newly inserted fitsTo
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve fitsTo ID: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "FitsTo added successfully!", "id": id})
}

func SetWineFitsTo(c *gin.Context) {
	wineID, err := strconv.Atoi(c.Query("wine_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wine_id"})
		return
	}
	fitsTo1, err := strconv.Atoi(c.Query("fitsTo1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fitsTo1"})
		return
	}
	fitsTo2, err := strconv.Atoi(c.Query("fitsTo2"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fitsTo2"})
		return
	}
	fitsTo3, err := strconv.Atoi(c.Query("fitsTo3"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fitsTo3"})
		return
	}

	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert or update the wine fitsTo
	query := `INSERT INTO Wine_FitsTo (wine_id, fitsTo_id_1, fitsTo_id_2, fitsTo_id_3)
	          VALUES (?, ?, ?, ?)
	          ON CONFLICT(wine_id) DO UPDATE SET
	          fitsTo_id_1=excluded.fitsTo_id_1, fitsTo_id_2=excluded.fitsTo_id_2, fitsTo_id_3=excluded.fitsTo_id_3`
	_, err = db.Exec(query, wineID, fitsTo1, fitsTo2, fitsTo3)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to set wine fitsTo: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wine fitsTo set successfully!"})
}
