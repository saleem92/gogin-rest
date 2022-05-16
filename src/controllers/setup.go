package controllers

import (
	"context"
	"go-rest/basics/models"
	"log"
	"net/http"
	"strconv"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func publishMessage(c *gin.Context) {

	client, err := dapr.NewClient()

	if err != nil {
		panic(err)
	}

	defer client.Close()

	ctx := context.Background()
	if err := client.PublishEvent(ctx, "pubsub", "orders", []byte(strconv.Itoa(10))); err != nil {
		panic(err)
	}

	log.Println("published data")
}

func SetupRoutes(router *gin.Engine) {
	router.GET("/", getAlbums)
	router.POST("/", postAlbum)
	router.GET("/:id", getAlbumById)
	router.PATCH("/:id", updateAlbum)
	router.DELETE("/:id", deleteAlbum)
	router.POST("/publish-message", publishMessage)
	router.GET("/health", checkHealth)
}

func checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Running"})
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum models.Album

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
	var album models.Album

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
