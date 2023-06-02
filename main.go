package main

import (
	"net/http"

	"golang-practice/pkg/baby"

	"github.com/gin-gonic/gin"
)

var babies = []baby.Baby{
	{Id: "1", FirstName: "Toto", LastName: "Tata", BirthDate: "2022-12-23"},
	{Id: "2", FirstName: "titi", LastName: "Titi", BirthDate: "2023-12-23"},
	{Id: "3", FirstName: "Tutu", LastName: "Tutu", BirthDate: "2024-12-23"},
}

func getBabies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, babies)
}

// func latencyMiddleware(c *gin.Context) {
//  t := time.Now()
//  c.Next()
//  latency := time.Since(t)
//  logger.Print("Latency: " + latency.String())
// }

// postAlbums adds an album from JSON received in the request body.
func addBaby(c *gin.Context) {
	var baby baby.Baby

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&baby); err != nil {
		return
	}

	// Add the new album to the slice.
	babies = append(babies, baby)
	c.IndentedJSON(http.StatusCreated, baby)
}

func getBabyById(c *gin.Context) {
	id := c.Param("id")

	for _, baby := range babies {
		if baby.Id == id {
			c.IndentedJSON(http.StatusOK, baby)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Baby not found"})
}

func buildRouter() *gin.Engine {
	router := gin.Default()
	//router.Use(latencyMiddleware)

	router.GET("/babies", getBabies)
	router.GET("/babies/:id", getBabyById)
	router.POST("/babies", addBaby)

	return router
}

func main() {
	var router = buildRouter()
	router.Run(":8080")
}
