package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FilterWine(c *gin.Context) {
	color := c.Query("color")
	fit := c.Query("fit")
	flavour := c.Query("flavour")
	price := c.Query("price")

	// Open the SQLite database file
	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Build the query
	query := `
	SELECT 
		Wine.id, 
		Wine.name, 
		Wine.year, 
		Wine.volume, 
		Wine.imageURL, 
		MIN(WineSupermarkets.price) AS cheapestPrice 
	FROM Wine 
	LEFT JOIN WineSupermarkets ON Wine.id = WineSupermarkets.wine_id 
	LEFT JOIN Wine_Flavour ON Wine.id = Wine_Flavour.wine_id
	LEFT JOIN Flavour ON Wine_Flavour.flavour_id_1 = Flavour.id OR Wine_Flavour.flavour_id_2 = Flavour.id OR Wine_Flavour.flavour_id_3 = Flavour.id
	LEFT JOIN Wine_FitsTo ON Wine.id = Wine_FitsTo.wine_id
	LEFT JOIN FitsTo ON Wine_FitsTo.fitsTo_id_1 = FitsTo.id OR Wine_FitsTo.fitsTo_id_2 = FitsTo.id OR Wine_FitsTo.fitsTo_id_3 = FitsTo.id
	`

	var filters []string

	// Filter by color
	if color != "" && color != "all" {
		filters = append(filters, "Wine.type = '"+color+"'")
	}

	// Filter by fit or flavour
	if fit != "" {
		filters = append(filters, "FitsTo.description = '"+fit+"'")
	} else if flavour != "" {
		filters = append(filters, "Flavour.name = '"+flavour+"'")
	}

	// Add filters to query
	if len(filters) > 0 {
		query += " WHERE " + filters[0]
		for _, filter := range filters[1:] {
			query += " AND " + filter
		}
	}

	query += " GROUP BY Wine.id"

	// Sort by price
	orderBy := " ORDER BY cheapestPrice ASC"
	if price == "1" {
		orderBy = " ORDER BY cheapestPrice DESC"
	}
	query += orderBy

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// Prepare the response
	var wines []struct {
		ID            int     `json:"id"`
		Name          string  `json:"name"`
		Year          int     `json:"year"`
		Volume        float64 `json:"volume"`
		ImageURL      string  `json:"imageURL"`
		CheapestPrice float64 `json:"cheapestPrice"`
	}

	for rows.Next() {
		var wine struct {
			ID            int             `json:"id"`
			Name          string          `json:"name"`
			Year          int             `json:"year"`
			Volume        float64         `json:"volume"`
			ImageURL      sql.NullString  `json:"imageURL"`
			CheapestPrice sql.NullFloat64 `json:"cheapestPrice"`
		}
		err = rows.Scan(&wine.ID, &wine.Name, &wine.Year, &wine.Volume, &wine.ImageURL, &wine.CheapestPrice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Handling NULL values
		var imageURL string
		if wine.ImageURL.Valid {
			imageURL = wine.ImageURL.String
		} else {
			imageURL = ""
		}

		var cheapestPrice float64
		if wine.CheapestPrice.Valid {
			cheapestPrice = wine.CheapestPrice.Float64
		} else {
			cheapestPrice = 0.0
		}

		wines = append(wines, struct {
			ID            int     `json:"id"`
			Name          string  `json:"name"`
			Year          int     `json:"year"`
			Volume        float64 `json:"volume"`
			ImageURL      string  `json:"imageURL"`
			CheapestPrice float64 `json:"cheapestPrice"`
		}{
			ID:            wine.ID,
			Name:          wine.Name,
			Year:          wine.Year,
			Volume:        wine.Volume,
			ImageURL:      imageURL,
			CheapestPrice: cheapestPrice,
		})
	}

	// Return the response
	c.JSON(http.StatusOK, wines)
}
