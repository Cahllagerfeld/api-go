package album

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Get all albums
// @Schemes
// @Description get all albums
// @Accept json
// @Produce json
// @Router /albums [get]
// @Tags Albums
// @Success 200 {array} Album Album with UUID
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// @Summary Create a new Album
// @Schemes
// @Accept json
// @Produce json
// @Router /albums [post]
// @Tags Albums
// @Param data body albumDTO true "DTO for Album creation"
// @Success 200 {object} Album UUID of new Album
func CreateAlbum(c *gin.Context) {
	var NewAlbumDTO albumDTO

	if err := c.BindJSON(&NewAlbumDTO); err != nil {
		return
	}

	NewAlbum := Album{
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
