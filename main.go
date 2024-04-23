package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID		string 	`json:"id"`
	Title 	string 	`json:"title"`
	Artist 	string 	`json:"artist"`
	Price	float64 `json:"price"`
}

//album slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Green Train", Artist: "Alonce Durre", Price: 165.99},
	{ID: "2", Title: "JeJe", Artist: "Genry Muccho", Price: 16.99},
	{ID: "3", Title: "Sarruh Vinny", Artist: "Pyatachoks", Price: 41.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

//getAlbums responds width the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}