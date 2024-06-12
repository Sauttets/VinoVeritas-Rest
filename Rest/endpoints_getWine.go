package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Wine struct {
	Name          string  `json:"name"`
	CheapestPrice float64 `json:"cheapest_price"`
	Volume        float64 `json:"volume"`
	Year          int     `json:"year"`
}

func GetWine(c *gin.Context) {
	id := c.Param("id")

	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `
	SELECT Wine.name, Wine.year, Wine.volume, MIN(WineSupermarkets.price) as cheapest_price
	FROM Wine
	JOIN WineSupermarkets ON Wine.id = WineSupermarkets.wine_id
	WHERE Wine.id = ?
	GROUP BY Wine.id, Wine.name, Wine.year, Wine.volume
	LIMIT 1`

	var wine Wine
	err = db.QueryRow(query, id).Scan(&wine.Name, &wine.Year, &wine.Volume, &wine.CheapestPrice)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Wine not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, wine)
}
