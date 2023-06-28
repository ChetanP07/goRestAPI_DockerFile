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
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {

	route := gin.Default()

	route.GET("/albums", getAlbums)
	route.POST("/albums", postAlbums)
	route.GET("/albums/:id", getAlbumByID)
	route.DELETE("/albums/:id", deleteAlbumByID)
	route.PUT("/albums/:id", updateAlbum)

	route.Run(":8000")

}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {

	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

func getAlbumByID(c *gin.Context) {

	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}

func deleteAlbumByID(c *gin.Context) {

	Id := c.Param("Id")
	for index, a := range albums {
		if a.ID == Id {
			albums = append(albums[:index], albums[index+1:]...)
			break
		}
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func updateAlbum(c *gin.Context) {
	Id := c.Param("Id")
	for index, a := range albums {
		if a.ID == Id {
			albums = append(albums[:index], albums[index+1:]...)
			var newAlbum album
			if err := c.BindJSON(&newAlbum); err != nil {
				return
			}
			albums = append(albums, newAlbum)
			c.IndentedJSON(http.StatusCreated, newAlbum)
		}
	}
}
