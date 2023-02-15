package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.

type vm struct {
	Name        string `json:"name"`
	Description string `json:"os-type"`
	OStype      string `json:"artist"`
	OSvariant   string `json:"os-variant"`
	ram         int    `json:"ram"`
	vcpus       int    `json:"vcpus"`
	diskpath    string `json:"diskpath"`
	bus         string `json:"bus"`
	graphics    string `json:"graphics"`
	cdrom       string `json:"cdrom"`
	network     string `json:"network"`
}

// albums slice to seed record album data..
var vmlists = []vm{
	//{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	//{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	//{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},

}

// getAlbums responds with the list of all albums as JSON.
func getvm(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, vmlists)
}

// postAlbums adds an album from JSON received in the request body.
func postvm(c *gin.Context) {
	var newVm vm

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newVm); err != nil {
		return
	}

	// Add the new album to the slice.
	vmlists = append(vmlists, newVm)
	c.IndentedJSON(http.StatusCreated, newVm)
}

func main() {
	router := gin.Default()
	router.GET("/vm", getvm)
	router.POST("/vm", postvm)

	router.Run("localhost:8080")
}
