package main

import (
	"github.com/dkallman13/Vicky3Optimizer/initial"
	"github.com/gin-gonic/gin"
)

func init() {
	initial.GetEnvVars()
	initial.SaveFileRead()
	initial.ConnectToDB()
}

func main() {
	
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on localhost
}
