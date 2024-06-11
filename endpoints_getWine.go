package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getWine(c *gin.Context) {
	// TODO: access database and search for wine with id

	c.JSON(http.StatusNotFound, gin.H{"error": "Wine not found"})
}

func getFullWine(c *gin.Context) {

	c.JSON(http.StatusNotFound, gin.H{"error": "Wine not found"})
}
