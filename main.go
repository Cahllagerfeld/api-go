package main

import (
	_ "github.com/cahllagerfeld/go-api/docs"
	"github.com/cahllagerfeld/go-api/internal/album"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Music API
// @version 1.0
// @description Dummy API as first go project

func main() {
	router := gin.Default()
	router.GET("/albums", album.GetAlbums)
	router.POST("/albums", album.CreateAlbum)
	router.DELETE("/albums/:id", album.DeleteByID)
	router.PUT("/albums/:id", album.UpdateByID)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":3000")
}
