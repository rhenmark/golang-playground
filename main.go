package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type response struct {
	Data 	interface{} `json:"data"`
	Status 	int			`json:"status"`
	Error  *string 	`json:"error,omitempty"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func formatResponse(data interface{}, status int, err *string) response {
	response := response{
		Data: data,
		Status: int(status),
	}
	if err != nil {
		response.Error = err
	}

	return response
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			// c.IndentedJSON(http.StatusOK, gin.H{"data": a, "status": http.StatusOK, "error": nil})
			var response = formatResponse(a, int(http.StatusOK), nil)
			c.IndentedJSON(http.StatusOK, response)
			return
		}
	}
	var msg *string
	tempMsg := "album not found"
	msg = &tempMsg
	var response = formatResponse(nil, int(http.StatusNotFound), msg)
	c.IndentedJSON(http.StatusNotFound, response)
}

func getPokemonList(c *gin.Context) {
	var api = "https://pokeapi.co/api/v2/pokemon/ditto"

	response, err := http.Get(api)
	if err != nil {
		var msg *string
		tempMsg := "Something went wrong"
		msg = &tempMsg
		var res = formatResponse(nil, http.StatusBadRequest, msg)
		c.IndentedJSON(http.StatusBadRequest, res)
		return
	}

	responseData, err := ioutil.ReadAll(response.Body)

	var res = formatResponse(string(responseData), http.StatusBadRequest, nil)
	c.IndentedJSON(http.StatusOK, res)

	fmt.Println(string(responseData))
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.GET("/pokemons", getPokemonList)
	router.Run("localhost:8080")
}
