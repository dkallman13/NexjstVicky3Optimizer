package main

import (
	"github.com/gin-gonic/gin"
	"github.com/dkallman13/Vicky3Optimizer/initial"
)
func init(){
	initial.GetEnvVars()
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