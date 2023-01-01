package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 54.5},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 44.35},
	{ID: "3", Title: "Sarah Magan", Artist: "Julian Juli", Price: 60.55},
}

// getAlbums responds with list of all album
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// get Album by id
func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	// loop over the list
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbum adds an album from JSON received
func postAlbum(c *gin.Context) {
	var newAlbum album
	// Call BindJSON to bind the received
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}

func main() {
	router := gin.Default()
	// get albums
	router.GET("/albums", getAlbums)
	// post album
	router.POST("/albums", postAlbum)
	// get album
	router.GET("/albums/:id", getAlbumById)
	// server
	router.Run("localhost:8080")

}
