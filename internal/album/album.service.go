package album

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func CreateAlbum(c *gin.Context) {
	var NewAlbumDTO albumDTO

	if err := c.BindJSON(&NewAlbumDTO); err != nil {
		return
	}

	NewAlbum := album{
		ID:     uuid.New().String(),
		Title:  NewAlbumDTO.Title,
		Artist: NewAlbumDTO.Artist,
		Price:  NewAlbumDTO.Price,
	}

	albums = append(albums, NewAlbum)
	c.IndentedJSON(http.StatusCreated, NewAlbum)
}

func GetByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
