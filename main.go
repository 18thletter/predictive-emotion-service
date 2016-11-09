package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Initialize the database
	initDb()

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

func initDb() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	checkErr(err, "Error opening database")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS emotions (
			id serial PRIMARY KEY,
			emotion character varying(255)
		)
	`)
	checkErr(err, "Error creating table")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS heartbeats (
			id serial PRIMARY KEY,
			start_time timestamp,
			end_time timestamp
		)
	`)
	checkErr(err, "Error creating table")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS datasets (
			id serial PRIMARY KEY,
			created_at timestamp,
			updated_at timestamp,
			emotion_id int REFERENCES emotions (id)
		)
	`)
	checkErr(err, "Error creating table")
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalf(msg + ": %q", err)
	}
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
