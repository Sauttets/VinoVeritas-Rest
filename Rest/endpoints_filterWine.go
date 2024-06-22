package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FilterWine(c *gin.Context) {
	color := c.Query("color")
	fit := c.Query("fit")
	flavour := c.Query("flavour")
	price := c.Query("price")
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

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
		Wine.country, 
		Wine.type, 
		Wine.description, 
		Wine.imageURL, 
		Wine.volume, 
		Wine.volAlc, 
		Wine.dryness, 
		Wine.acidity, 
		Wine.tanninLevel
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
	orderBy := " ORDER BY MIN(WineSupermarkets.price) ASC"
	if price == "1" {
		orderBy = " ORDER BY MIN(WineSupermarkets.price) DESC"
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
	var wines []FullWine

	for rows.Next() {
		var wine FullWine
		err = rows.Scan(&wine.ID, &wine.Name, &wine.Year, &wine.Country, &wine.Type, &wine.Description, &wine.ImageURL, &wine.Volume, &wine.VolAlc, &wine.Dryness, &wine.Acidity, &wine.TanninLevel)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Check if the wine is in the user's favorite list
		favoriteQuery := `SELECT COUNT(1) FROM FavoriteWines WHERE user_id = ? AND wine_id = ?`
		var count int
		err = db.QueryRow(favoriteQuery, userID, wine.ID).Scan(&count)
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
		err = db.QueryRow(flavoursQuery, wine.ID).Scan(&flavour1, &flavour2, &flavour3)
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
		err = db.QueryRow(fitsToQuery, wine.ID).Scan(&fitsTo1, &fitsTo2, &fitsTo3)
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
		sRows, err := db.Query(supermarketsQuery, wine.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer sRows.Close()

		var supermarkets []Supermarket
		for sRows.Next() {
			var supermarket Supermarket
			if err := sRows.Scan(&supermarket.Name, &supermarket.Street, &supermarket.PostalCode, &supermarket.City, &supermarket.HouseNumber, &supermarket.Price); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			supermarkets = append(supermarkets, supermarket)
		}
		wine.Supermarkets = supermarkets

		wines = append(wines, wine)
	}

	// Return the response
	c.JSON(http.StatusOK, wines)
}
