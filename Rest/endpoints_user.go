package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUser(c *gin.Context) {
	username := c.Param("username")

	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert the new user into the database
	query := `INSERT INTO Users (username) VALUES (?)`
	result, err := db.Exec(query, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the ID of the newly inserted user
	userID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": userID, "username": username})
}

func UpdateUser(c *gin.Context) {
	username := c.Param("username")
	id := c.Param("id")

	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Update the username of the user with the given id
	query := `UPDATE Users SET username = ? WHERE id = ?`
	_, err = db.Exec(query, username, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
