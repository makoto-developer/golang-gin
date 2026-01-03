package models

// Album represents an album entity
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	Tax    float32 `json:"tax"`
}

// In-memory album store
var Albums = []Album{
	{ID: "1", Title: "Hammerhead", Artist: "THE OFFSPRING", Price: 25.05, Tax: 0.1},
	{ID: "2", Title: "Shake It Off", Artist: "Taylor Swift", Price: 23.14, Tax: 0.1},
	{ID: "3", Title: "mysterious love", Artist: "Miho Komatsu", Price: 18.88, Tax: 0.1},
}
