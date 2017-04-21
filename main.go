package main

import (
	"log"
	"net/http"
	"os"
  "fmt"
	"github.com/gin-gonic/gin"
)

func (message *Message) filter() (response []Record) {
	response = make([]Record, 0, len(message.Payload))
	for _, show := range message.Payload {
    if show.Drm && show.EpisodeCount > 0 {
			response = append(response, Record{show.Img.ShowImage, show.Slug, show.Title})
		}
	}
	return
}

func action(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
		  c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Could not decode request: " + r.(string)})
		}
	}()
	var msg Message
	c.BindJSON(&msg)
  fmt.Println(msg)
	if msg.Payload == nil {
		panic("JSON parsing failed")
	}
	response := msg.filter()
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.POST("/", action)

	router.Run(":" + port)
}
