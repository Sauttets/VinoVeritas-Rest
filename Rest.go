package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Wine struct {
    ID int				`json:"id"`			
    Name string			`json:"name"`		
    Year int			`json:"year"`		
    TasteID int 		`json:"taste_id"`	
    FitID int			`json:"fit_id"`		
    Volume float64		`json:"volume"`		
    VolAlc float64		`json:"vol_alc"`
	description string	`json:"description"`
    Shops []Shop		`json:"shops"`		
}

type Shop struct {
    Name string			`json:"name"`
    City string			`json:"city"`
    CityCode int		`json:"city_code"`
    Street string		`json:"street"`
    HouseNumber int		`json:"house_number"`
    Price int			`json:"price"`
}

var w1 = Wine{
    ID: 1,
    Name: "Chardonnay",
    Year: 2018,
    TasteID: 1,
    FitID: 1,
    Volume: 0.75,
    VolAlc: 13.5,
	description: "A dry white wine with a fruity taste",
    Shops: []Shop{
        {
            Name: "Aldi",
            City: "Munich",
            CityCode: 80331,
            Street: "Kaufingerstraße",
            HouseNumber: 5,
            Price: 5,
        },
        {
            Name: "Rewe",
            City: "Munich",
            CityCode: 80331,
            Street: "Marienplatz",
            HouseNumber: 1,
            Price: 6,
        },
    },
}

var wineList = []Wine{w1}

func main() {
	router := gin.Default()
	router.GET("/getWine/:id", getWine)
	router.GET("/getFullWine/:id", getFullWine)
	router.GET("/getFavList/:key", getFavList)
	router.POST("/updateFavList/:key/:username/:ID", updateFavList)
	router.GET("/getWineFactOTD", getWineFactOTD)
	router.GET("/getWineFacts", getwineFacts)
	router.Run("localhost:8080")
}

func getFullWine(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }
	//TODO: access database and search for wine with id 
    for _, wine := range wineList {
		print("searching for wine with id: ", id, " found wine with id: ", wine.ID, "\n")
        if wine.ID == id {
            c.JSON(http.StatusOK, wine)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Wine not found"})
}

func getWine(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }
	// TODO: access database and search for wine with id
    for _, wine := range wineList {
        if wine.ID == id {
            // Create a map with only the desired fields
            wineResponse := map[string]interface{}{
                "name":  wine.Name,
                "year":  wine.Year,
                "price": wine.Shops[0].Price, // Assuming the price is from the first shop
            }
            c.JSON(http.StatusOK, wineResponse)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Wine not found"})
}

func updateFavList(c *gin.Context) {
	//key := c.Param("key")
	//username := c.Param("username")
	//wineID := c.Param("ID")
	
	//TODO: access database and update favorite list at key with wineID
}

func getFavList(c *gin.Context) {
	//key := c.Param("key")

	//TODO: access database and get favorite list and Usename at key 
}

func getWineFactOTD(c *gin.Context) {
	
	// TODO: access database and get wine fact OTD
}

func getwineFacts(c *gin.Context) {

	// TODO: access database and get wine facts
}

