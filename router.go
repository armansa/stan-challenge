package main

import "github.com/gin-gonic/gin"

func SetupRouting(port string) {
  	router := gin.New()
  	router.Use(gin.Logger())
  	router.LoadHTMLGlob("templates/*.tmpl.html")
  	router.Static("/static", "static")

  	router.POST("/", MainAction)

  	router.Run(":" + port)
}
