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
	router.POST("/AddToFavList", AddToFavList)
	router.POST("/deleteFromFavList", DeleteFromFavList)

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

	router.GET("/filterWine", FilterWine)

	router.Run("localhost:8083")
}

/*

export CC=x86_64-unknown-linux-gnu-gcc
CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -o VinoVeritasRest

unset CC;unset GOARCH;unset GOOS;unset CGO_ENABLED


*/
