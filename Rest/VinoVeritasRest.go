package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	SetupDB()
	router := gin.Default()
	//id
	router.GET("/getWine", GetWine)
	//id
	router.GET("/getFullWine", GetFullWine)
	router.GET("/getWineFactOTD", GetWineFactOTD)

	//user_id
	router.GET("/getFavList", GetFavList)
	//user_id & username & wine_id
	router.POST("/updateFavList", AddToFavList)
	//user_id & username & wine_id
	router.POST("/deleteFavList", DeleteFromFavList)

	//username
	router.POST("/newUser", NewUser)
	///username & id
	router.POST("/updateUser", UpdateUser)

	///name & year & country & type & description & imageURL & volume & volAlc
	router.POST("/addWine", authRequired(), AddWine)
	///:wine_id/:flavour1/:flavour2 & flavour3
	router.POST("/setWineFlavour", authRequired(), SetWineFlavour)
	//wine_id & fitsTo1 & fitsTo2 & fitsTo3
	router.POST("/setWineFit", authRequired(), SetWineFitsTo)

	//add Flavour and FitsTo
	//name
	router.POST("/addFlavour", authRequired(), AddFlavour)
	// name
	router.POST("/addFitsTo", authRequired(), AddFitsTo)

	//name & street & postal_code & city & houseNumber
	router.POST("/addSupermarket", authRequired(), AddSupermarket)
	//id & supermarket_id & price
	router.POST("/setWineSupermarket", authRequired(), setWineSupermarket)

	router.Run("localhost:8083")
}
