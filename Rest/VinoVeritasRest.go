package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	SetupDB()
	router := gin.Default()
	router.GET("/getWines", GetWines)
	router.GET("/getWineFactOTD", GetWineFactOTD)
	router.POST("/AddToFavList", AddToFavList)
	router.POST("/deleteFromFavList", DeleteFromFavList)
	router.POST("/newUser", NewUser)
	router.POST("/updateUser", UpdateUser)
	router.POST("/addWine", authRequired(), AddWine)
	router.POST("/addSupermarket", authRequired(), AddSupermarket)
	router.POST("/addFlavour", authRequired(), AddFlavour)
	router.POST("/addFitsTo", authRequired(), AddFitsTo)
	router.POST("/setWineFlavour", authRequired(), SetWineFlavour)
	router.POST("/setWineFit", authRequired(), SetWineFitsTo)
	router.POST("/setWineSupermarket", authRequired(), setWineSupermarket)
	router.Run("localhost:8083")
}

/*

export CC=x86_64-unknown-linux-gnu-gcc
CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -o VinoVeritasRest

unset CC;unset GOARCH;unset GOOS;unset CGO_ENABLED


*/
