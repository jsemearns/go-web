package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	// routes

	// GET
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)

	// POST
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(context *gin.Context) {
	var newAlbum album

	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(context *gin.Context) {
	id := context.Param("id")

	for _, album := range albums {
		if album.ID == id {
			context.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
