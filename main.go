package main

import (
	"net/http"
	"strings"

	"github.com/dkallman13/Vicky3Optimizer/initial"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func init() {
	initial.GetEnvVars()
	initial.SaveFileLocSetter()
	initial.ConnectToDB()
	initial.SaveFileLister()
}

func createRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	
	r.AddFromFiles("index", "templates/base/base.html", "templates/index.html")
	r.AddFromFiles("/save", "templates/base/base.html", "templates/save/save.html")
	return r
}

func main() {
	router := gin.Default()
	router.HTMLRender = createRenderer()

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index" ,  gin.H{
			"title": "home",
			"savefiles": initial.SaveFiles,
		})
	})
	for _, savFile := range(initial.SaveFiles){
		var savRouteBuilder strings.Builder
		savRouteBuilder.WriteString("save/")
		savRouteBuilder.WriteString(savFile)
		router.GET(savRouteBuilder.String() , func(c *gin.Context) {
		c.HTML(http.StatusOK, "/save" ,  gin.H{
			"title": savFile,
			"firstline" : initial.TokenizeSave(savFile),
		})
	})
	}
	
	router.Run() // listen and serve on localhost
}
