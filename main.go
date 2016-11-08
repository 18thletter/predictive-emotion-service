package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	// Serve the docs page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	// Group v1 API resources
	v1 := router.Group("/v1")
	{
		v1.GET("/datasets", GetAllDatasets)
		v1.POST("/datasets", CreateDataset)
		v1.POST("/datasets/:dataset_id", GetDataset)
		v1.POST("/heartbeats", CreateHeartbeat)
		v1.POST("/emotions", CreateEmotion)
		v1.GET("/emotions", GetAllEmotions)
		v1.POST("/correct", CorrectDataset)
		v1.GET("/predict", GetPrediction)
	}

	router.Run(":" + port)
}

func GetAllDatasets(c *gin.Context) {
	c.JSON(http.StatusOK, "{}")
}

func CreateDataset(c *gin.Context) {
	c.JSON(http.StatusOK, "{}")
}

func GetDataset(c *gin.Context) {
	c.JSON(http.StatusOK, "{}")
}

func CreateHeartbeat(c *gin.Context) {
	c.JSON(http.StatusOK, "{}")
}

func CreateEmotion(c *gin.Context) {
	c.JSON(http.StatusOK, "{}")
}

func GetAllEmotions(c *gin.Context) {
	c.JSON(http.StatusOK, "{}")
}

func CorrectDataset(c *gin.Context) {
	c.JSON(http.StatusOK, "{}")
}

func GetPrediction(c *gin.Context) {
	c.JSON(http.StatusOK, "{}")
}
