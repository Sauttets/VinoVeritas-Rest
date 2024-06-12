package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddFlavour(c *gin.Context) {
	name := c.Query("flavour")

	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert new flavour entry
	query := `INSERT INTO Flavour (name) VALUES (?)`
	result, err := db.Exec(query, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add flavour: %v", err)})
		return
	}

	// Get the ID of the newly inserted flavour
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve flavour ID: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flavour added successfully!", "id": id})
}

func SetWineFlavour(c *gin.Context) {
	wineID, err := strconv.Atoi(c.Query("wine_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wine_id"})
		return
	}
	flavour1, err := strconv.Atoi(c.Query("flavour1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid flavour1"})
		return
	}
	flavour2, err := strconv.Atoi(c.Query("flavour2"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid flavour2"})
		return
	}
	flavour3, err := strconv.Atoi(c.Query("flavour3"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid flavour3"})
		return
	}

	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert or update the wine flavours
	query := `INSERT INTO Wine_Flavour (wine_id, flavour_id_1, flavour_id_2, flavour_id_3)
	          VALUES (?, ?, ?, ?)
	          ON CONFLICT(wine_id) DO UPDATE SET
	          flavour_id_1=excluded.flavour_id_1, flavour_id_2=excluded.flavour_id_2, flavour_id_3=excluded.flavour_id_3`
	_, err = db.Exec(query, wineID, flavour1, flavour2, flavour3)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to set wine flavours: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wine flavours set successfully!"})
}
