package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FullWine struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Year         int           `json:"year"`
	Country      string        `json:"country"`
	Type         string        `json:"type"`
	Description  string        `json:"description"`
	ImageURL     string        `json:"imageURL"`
	Volume       float64       `json:"volume"`
	VolAlc       float64       `json:"volAlc"`
	IsLiked      bool          `json:"isLiked"`
	Dryness      float64       `json:"dryness"`
	Acidity      float64       `json:"acidity"`
	TanninLevel  float64       `json:"tanninLevel"`
	Flavours     []string      `json:"flavours"`
	FitsTo       []string      `json:"fitsTo"`
	Supermarkets []Supermarket `json:"supermarkets"`
}

type Supermarket struct {
	Name        string  `json:"name"`
	Street      string  `json:"street"`
	PostalCode  string  `json:"postal_code"`
	City        string  `json:"city"`
	HouseNumber string  `json:"house_number"`
	Price       float64 `json:"price"`
}

func GetFullWine(c *gin.Context) {
	wineID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wine ID"})
		return
	}

	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Fetch the basic wine information
	var wine FullWine
	query := `
	SELECT id, name, year, country, type, description, imageURL, volume, volAlc, dryness, acidity, tanninLevel
	FROM Wine
	WHERE id = ?`
	err = db.QueryRow(query, wineID).Scan(&wine.ID, &wine.Name, &wine.Year, &wine.Country, &wine.Type, &wine.Description, &wine.ImageURL, &wine.Volume, &wine.VolAlc, &wine.Dryness, &wine.Acidity, &wine.TanninLevel)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Wine not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Check if the wine is in the user's favorite list
	favoriteQuery := `SELECT COUNT(1) FROM FavoriteWines WHERE user_id = ? AND wine_id = ?`
	var count int
	err = db.QueryRow(favoriteQuery, userID, wineID).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	wine.IsLiked = count > 0

	// Fetch the flavours of the wine
	flavoursQuery := `
	SELECT f1.name, f2.name, f3.name
	FROM Wine_Flavour
	LEFT JOIN Flavour f1 ON Wine_Flavour.flavour_id_1 = f1.id
	LEFT JOIN Flavour f2 ON Wine_Flavour.flavour_id_2 = f2.id
	LEFT JOIN Flavour f3 ON Wine_Flavour.flavour_id_3 = f3.id
	WHERE Wine_Flavour.wine_id = ?`
	var flavour1, flavour2, flavour3 sql.NullString
	err = db.QueryRow(flavoursQuery, wineID).Scan(&flavour1, &flavour2, &flavour3)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if flavour1.Valid {
		wine.Flavours = append(wine.Flavours, flavour1.String)
	}
	if flavour2.Valid {
		wine.Flavours = append(wine.Flavours, flavour2.String)
	}
	if flavour3.Valid {
		wine.Flavours = append(wine.Flavours, flavour3.String)
	}

	// Fetch the fits to of the wine
	fitsToQuery := `
	SELECT ft1.description, ft2.description, ft3.description
	FROM Wine_FitsTo
	LEFT JOIN FitsTo ft1 ON Wine_FitsTo.fitsTo_id_1 = ft1.id
	LEFT JOIN FitsTo ft2 ON Wine_FitsTo.fitsTo_id_2 = ft2.id
	LEFT JOIN FitsTo ft3 ON Wine_FitsTo.fitsTo_id_3 = ft3.id
	WHERE Wine_FitsTo.wine_id = ?`
	var fitsTo1, fitsTo2, fitsTo3 sql.NullString
	err = db.QueryRow(fitsToQuery, wineID).Scan(&fitsTo1, &fitsTo2, &fitsTo3)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if fitsTo1.Valid {
		wine.FitsTo = append(wine.FitsTo, fitsTo1.String)
	}
	if fitsTo2.Valid {
		wine.FitsTo = append(wine.FitsTo, fitsTo2.String)
	}
	if fitsTo3.Valid {
		wine.FitsTo = append(wine.FitsTo, fitsTo3.String)
	}

	// Fetch the supermarkets carrying the wine and their prices
	supermarketsQuery := `
	SELECT Supermarkets.name, Supermarkets.street, Supermarkets.postal_code, Supermarkets.city, Supermarkets.house_number, WineSupermarkets.price
	FROM WineSupermarkets
	JOIN Supermarkets ON WineSupermarkets.supermarket_id = Supermarkets.id
	WHERE WineSupermarkets.wine_id = ?`
	rows, err := db.Query(supermarketsQuery, wineID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var supermarkets []Supermarket
	for rows.Next() {
		var supermarket Supermarket
		if err := rows.Scan(&supermarket.Name, &supermarket.Street, &supermarket.PostalCode, &supermarket.City, &supermarket.HouseNumber, &supermarket.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		supermarkets = append(supermarkets, supermarket)
	}
	wine.Supermarkets = supermarkets

	c.JSON(http.StatusOK, wine)
}
