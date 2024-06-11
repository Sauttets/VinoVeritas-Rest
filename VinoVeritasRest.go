package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/getWine/:id", getWine)
	router.GET("/getFullWine/:id", getFullWine)
	router.GET("/getFavList/:key", getFavList)
	router.POST("/updateFavList/:key/:username/:ID", updateFavList)
	router.GET("/getWineFactOTD", getWineFactOTD)
	router.GET("/getWineFacts", getwineFacts)
	router.POST("/newUser/:username", newUser)
	router.Run("localhost:8080")
}
