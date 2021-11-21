package album

import "github.com/google/uuid"

var albums = []Album{
	{
		ID:     uuid.New().String(),
		Title:  "ADSF",
		Artist: "Baui",
		Price:  40.99,
	},
}

type albumDTO struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}
