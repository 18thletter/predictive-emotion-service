package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB
var inMemoryState string

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
	log.Println("Connected to PostreSQL database.")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS emotions (
			emotion character varying(255) PRIMARY KEY NOT NULL
		)
	`)
	checkErr(err, "Error creating table")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS heartbeats (
			id serial PRIMARY KEY NOT NULL,
			start_time timestamp NOT NULL,
			end_time timestamp NOT NULL
		)
	`)
	checkErr(err, "Error creating table")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS datasets (
			id serial PRIMARY KEY NOT NULL,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			emotion character varying(255) REFERENCES emotions (emotion) NOT NULL
		)
	`)
	checkErr(err, "Error creating table")
	inMemoryState = ""
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalf(msg + ": %q", err)
	}
}

type Dataset struct {
	Emotion string `form:"emotion" json:"emotion" binding:"required"`
}

type Emotion struct {
	Emotion string `form:"emotion" json:"emotion" binding:"required"`
}

func GetAllDatasets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func CreateDataset(c *gin.Context) {
	var json Dataset
	var datasetId int
	var emotion string
	var createdAt time.Time
	var updatedAt time.Time

	if c.BindJSON(&json) == nil {
		now := time.Now()
		row := db.QueryRow(`INSERT INTO datasets(created_at, updated_at, emotion)
			VALUES ($1,$2,$3) RETURNING *`, now, now, json.Emotion)
		row.Scan(&datasetId, &createdAt, &updatedAt, &emotion)
	}

	c.JSON(http.StatusOK, gin.H{
		"id": datasetId,
		"createdAt": createdAt,
		"updatedAt": updatedAt,
		"emotion": emotion,
	})
}

func GetDataset(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func CreateHeartbeat(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func CreateEmotion(c *gin.Context) {
	var json Emotion
	if c.BindJSON(&json) == nil {
		db.QueryRow("INSERT INTO emotions(emotion) VALUES($1)", json.Emotion)
	}
	c.JSON(http.StatusOK, gin.H{"emotion": json.Emotion})
}

func GetAllEmotions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func CorrectDataset(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func GetPrediction(c *gin.Context) {
	emotion := c.Query("e")
	if (emotion == "" && inMemoryState == "") {
		emotion = "happy"
	} else if (emotion == "" && inMemoryState != "") {
		emotion = inMemoryState
	}
	inMemoryState = emotion
	c.JSON(http.StatusOK, gin.H{"emotion": emotion})
}
