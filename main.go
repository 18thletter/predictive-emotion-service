package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	// Serve the docs page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	// Database migrations (heroku)
	router.POST("/migrate", MigrateFunc)

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

func MigrateFunc(c *gin.Context) {

	if _, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS emotions (
				id serial PRIMARY KEY,
				emotion character varying(255)
			)
		`); err != nil {
		c.JSON(http.StatusInternalServerError,
			fmt.Sprintf("Error creating database table: %q", err))
		return
	}
	//
	// if _, err := db.Exec(`
	// 		CREATE TABLE IF NOT EXISTS heartbeats (
	// 			id serial PRIMARY KEY,
	// 			start_time timestamp,
	// 			end_time timestamp
	// 		)
	// 	`); err != nil {
	// 	c.JSON(http.StatusInternalServerError,
	// 		fmt.Sprintf("Error creating database table: %q", err))
	// 	return
	// }
	//
	// if _, err := db.Exec(`
	// 		CREATE TABLE IF NOT EXISTS datasets (
	// 			id serial PRIMARY KEY,
	// 			created_at timestamp,
	// 			updated_at timestamp,
	// 			emotion_id int REFERENCES emotions (id)
	// 		)
	// 	`); err != nil {
	// 	c.JSON(http.StatusInternalServerError,
	// 		fmt.Sprintf("Error creating database table: %q", err))
	// 	return
	// }

	// if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp)"); err != nil {
	// 	c.String(http.StatusInternalServerError,
	// 		fmt.Sprintf("Error creating database table: %q", err))
	// 	return
	// }
	//
	// if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
	// 	c.String(http.StatusInternalServerError,
	// 		fmt.Sprintf("Error incrementing tick: %q", err))
	// 	return
	// }
	//
	// rows, err := db.Query("SELECT tick FROM ticks")
	// if err != nil {
	// 	c.String(http.StatusInternalServerError,
	// 		fmt.Sprintf("Error reading ticks: %q", err))
	// 	return
	// }
	//
	// defer rows.Close()
	// for rows.Next() {
	// 	var tick time.Time
	// 	if err := rows.Scan(&tick); err != nil {
	// 		c.String(http.StatusInternalServerError,
	// 			fmt.Sprintf("Error scanning ticks: %q", err))
	// 		return
	// 	}
	// }
	c.JSON(http.StatusOK, "{}")
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
