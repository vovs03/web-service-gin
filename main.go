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
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

//getAlbums responds width the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON recieved in the request bidy.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BuildJSON to bind the recieved JSON
	// to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

/* getAlbumById locates the album whose ID value matches 
the id parameter sent the client, then returns that album as a response*/
func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	/* Loop over the list of albums, looking for
	an album whose ID value matches the parameter */
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

