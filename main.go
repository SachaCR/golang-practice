package main

import (
	"golang-practice/pkg/children"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var childrenList = children.ChildrenList

func getAllChildren(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, childrenList)
}

func addChild(c *gin.Context) {
	var children children.Children

	err := c.BindJSON(&children)

	if err != nil {
		return
	}

	// Add the new album to the slice.
	childrenList = append(childrenList, children)

	c.IndentedJSON(http.StatusCreated, children)
}

func getChildById(c *gin.Context) {
	id := c.Param("id")

	for _, children := range childrenList {
		if children.Id == id {
			c.IndentedJSON(http.StatusOK, children)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Child not found"})
}

func getChildTimeline(c *gin.Context) {

	c.IndentedJSON(http.StatusNotImplemented, gin.H{"message": "This feature is under development"})
}

func buildRouter(env string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	if env != "test" {
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
	}

	router.GET("/children", getAllChildren)
	router.GET("/children/:id", getChildById)
	router.GET("/children/:id/timeline/:date", getChildTimeline)
	router.POST("/children", addChild)

	return router
}

func main() {
	environment := os.Getenv("GO_ENV")

	var router = buildRouter(environment)
	router.Run(":8080")
}
