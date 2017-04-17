package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

type Episode struct {
  Channel string `json:"channel" binding:"required"`
  ChannelLogo string `json:"channelLogo" binding:"required"`
  Date string `json:"date" binding:"required"`
  Html string `json:"html" binding:"required"`
  Url string `json:"url" binding:"required"`
}

type Season struct {
  Slug string `json:"slug" binding:"required"`
}

type Image struct {
	ShowImage string `json:"showImage" binding:"required"`
}

type Show struct {
	Country string `json:"country" binding:"required"`
	Description string `json:"description" binding:"required"`
	Drm bool `json:"drm" binding:"required"`
	EpisodeCount int `json:"episodeCount" binding:"required"`
	Genre string `json:"genre" binding:"required"`
	Img Image `json:"image" binding:"required"`
	Language string `json:"language" binding:"required"`
	NextEpisode Episode `json:"nextEpisode" binding:"required"`
	PrimaryColour string `json:"primaryColour" binding:"required"`
	Seasons []Season `json:"seasons" binding:"required"`
	Slug string `json:"slug" binding:"required"`
	Title string `json:"title" binding:"required"`
	TvChannel string `json:"tvChannel" binding:"required"`
}

type Message struct {
	Payload []Show `json:"payload" binding:"required"`
	Skip int `json:"skip" binding:"required"`
	Take int `json:"take" binding:"required"`
	TotalRecords int `json:"totalRecords" binding:"required"`
}

type Record struct {
	Image string `json:"image" binding:"required"`
	Slug string `json:"slug" binding:"required"`
	Title string `json:"title" binding:"required"`
}

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
		  c.IndentedJSON(http.StatusBadRequest, gin.H{"error": r})
		}
	}()
	var msg Message
	c.BindJSON(&msg)
	if msg.Payload == nil {
		panic("action payload has invalid format")
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
