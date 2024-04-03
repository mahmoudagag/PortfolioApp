package main

import (
	// "net/http"

	// "github.com/gin-gonic/contrib/static"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/VisualizeSearchingAlogrithms", "./Visualize-Search-Algorithms")
	router.Static("/VisualizeSortingAlogrithms", "./Visualize-Sorting-Algorithms")

	router.Run("localhost:8080")
}
