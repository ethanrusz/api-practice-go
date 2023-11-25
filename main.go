package main

import (
	"api-practice/models"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	// API routing gubbins
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/album/:id", getAlbumByID)
	router.POST("/albums", addAlbum)

	err := router.Run("localhost:8080")

	if err != nil {
		panic(err.Error())
	}
}

func getAlbums(c *gin.Context) {
	albums := models.GetAlbums()

	if albums == nil || len(albums) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, albums)
	}
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album := models.GetAlbumByID(id)

	if album == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		fmt.Println(album)
		c.IndentedJSON(http.StatusOK, album)
	}
}

func addAlbum(c *gin.Context) {
	var album models.Album

	if err := c.BindJSON(&album); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddAlbum(album)
		c.IndentedJSON(http.StatusCreated, album)
	}
}
