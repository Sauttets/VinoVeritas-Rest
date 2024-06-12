package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddToFavList(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}
	wineID, err := strconv.Atoi(c.Query("wine_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wine_id"})
		return
	}

	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert new favorite wine entry
	query := `INSERT INTO FavoriteWines (user_id, wine_id) VALUES (?, ?)`
	_, err = db.Exec(query, userID, wineID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add to favorite list: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wine added to favorite list successfully!"})
}

func DeleteFromFavList(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}
	wineID, err := strconv.Atoi(c.Query("wine_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wine_id"})
		return
	}

	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Delete favorite wine entry
	query := `DELETE FROM FavoriteWines WHERE user_id = ? AND wine_id = ?`
	_, err = db.Exec(query, userID, wineID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete from favorite list: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wine removed from favorite list successfully!"})
}

func GetFavList(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Retrieve favorite wine IDs for the user
	query := `SELECT wine_id FROM FavoriteWines WHERE user_id = ?`
	rows, err := db.Query(query, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve favorite list: %v", err)})
		return
	}
	defer rows.Close()

	var favList []int
	for rows.Next() {
		var wineID int
		if err := rows.Scan(&wineID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to scan wine_id: %v", err)})
			return
		}
		favList = append(favList, wineID)
	}

	c.JSON(http.StatusOK, gin.H{"favorite_wines": favList})
}
