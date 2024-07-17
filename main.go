package main

import (
	"net/http"
	"github.com/gin-contrib/multitemplate"
	"github.com/dkallman13/Vicky3Optimizer/initial"
	"github.com/gin-gonic/gin"
)

func init() {
	initial.GetEnvVars()
	initial.SaveFileLocSetter()
	initial.ConnectToDB()
}

func createRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	
	r.AddFromFiles("index", "templates/base/base.html", "templates/index.html")
	return r
}

func main() {
	router := gin.Default()
	router.HTMLRender = createRenderer()

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index" ,  gin.H{
			"title": "home",
			"savefilelocation": initial.SaveFileLoc,
		})
	})
	router.Run() // listen and serve on localhost
}
