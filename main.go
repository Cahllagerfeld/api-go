package main

import (
	_ "github.com/cahllagerfeld/go-api/docs"
	"github.com/cahllagerfeld/go-api/internal/album"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	router.GET("/albums", album.GetAlbums)
	router.POST("/albums", album.CreateAlbum)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":3000")
}
