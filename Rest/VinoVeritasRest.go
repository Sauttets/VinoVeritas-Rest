package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	SetupDB()

	router := gin.Default()
	router.GET("/getWine/:id", GetWine)
	router.GET("/getFullWine/:id", GetFullWine)
	router.GET("/getWineFactOTD", GetWineFactOTD)

	router.GET("/getFavList/:user_id", GetFavList)
	router.POST("/updateFavList/:user_id/:username/:wine_id", AddToFavList)
	router.POST("/deleteFavList/:user_id/:username/:wine_id", DeleteFromFavList)

	router.POST("/newUser/:username", NewUser)
	router.POST("/updateUser/:username/:id", UpdateUser)

	router.POST("/addWineFactOTD/:name/:year/:country/:type/:description/:imageURL/:volume/:volAlc", authRequired(), AddWine)
	router.POST("/setWineFlavour/:wine_id/:flavour1/:flavour2/:flavour3", authRequired(), SetWineFlavour)
	router.POST("/setWineFitsTo/:wine_id/:fitsTo1/:fitsTo2/:fitsTo3", authRequired(), SetWineFitsTo)

	//add Flavour and FitsTo
	router.POST("/addFlavour/:name", authRequired(), AddFlavour)
	router.POST("/addFitsTo/:name", authRequired(), AddFitsTo)

	router.POST("/addSupermarket/:name/:street/:postal_code/:city/:houseNumber", authRequired(), AddSupermarket)
	router.POST("/setWineSupermarket/:id/:supermarket_id/:price", authRequired(), setWineSupermarket)

	router.Run("localhost:8083")
}
