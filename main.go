package main

import (
	"structure/api"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*") //loads all gtml static pages
	router.GET("/", api.Home)
	router.GET("/all", api.All)
	router.POST("/insert", api.Insert)
	router.POST("/delete", api.Delete)
	//log.Fatal("Server starting")
	router.Run()

}
