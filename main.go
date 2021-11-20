package main

import (
	"github.com/cahllagerfeld/go-api/internal/album"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", album.GetAlbums)
	router.POST("/albums", album.CreateAlbum)

	router.Run(":3000")
}
