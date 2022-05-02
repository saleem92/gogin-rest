package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id" binding:"required"`
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	gin.DisableConsoleColor()

	// Used to add file logger
	// file, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(file)

	router := gin.Default()
	router.GET("/", getAlbums)
	router.POST("/", postAlbum)
	router.GET("/:id", getAlbumById)
	router.PATCH("/:id", updateAlbum)
	router.DELETE("/:id", deleteAlbum)

	router.Run()
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	albums = append(albums, newAlbum)

	c.JSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	var album album

	if err := c.BindJSON(&album); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, a := range albums {
		if a.ID == id {
			albums[i] = album

			c.JSON(http.StatusOK, album)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)

			c.JSON(http.StatusOK, gin.H{"message": "Album Deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
