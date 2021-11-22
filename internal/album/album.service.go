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
// @Param Body body albumDTO true "DTO for Album creation"
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

// @Summary Delete a album Album
// @Schemes
// @Accept json
// @Router /albums/{id}  [delete]
// @Tags Albums
// @Param id path string true "ID of Item"
// @Success 204
func DeleteByID(c *gin.Context) {
	id := c.Param("id")
	k := 0
	for _, album := range albums {
		if album.ID == id {
			albums[k] = album
			k++
		}
	}
	albums = albums[:k]
}

// @Summary Update a album
// @Schemes
// @Accept json
// @Router /albums/{id}  [put]
// @Tags Albums
// @Param id path string true "ID of Item"
// @Param Body body albumDTO true "DTO for update the album"
// @Success 200
func UpdateByID(c *gin.Context) {
	id := c.Param("id")
	var NewAlbumDTO albumDTO

	if err := c.BindJSON(&NewAlbumDTO); err != nil {
		return
	}

	for _, album := range albums {
		if album.ID == id {
			album.ID = album.ID
			album.Artist = NewAlbumDTO.Artist
			album.Price = NewAlbumDTO.Price
			album.Title = NewAlbumDTO.Title
			c.IndentedJSON(http.StatusCreated, album)
		}
	}
}
