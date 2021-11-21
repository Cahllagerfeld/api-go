package album

var albums = []album{}

type albumDTO struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}
