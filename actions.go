package main

import (
	"net/http"
  "fmt"
	"github.com/gin-gonic/gin"
)

func MainAction(c *gin.Context) {
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
	response := msg.Filter()
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}
