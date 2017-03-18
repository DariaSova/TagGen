package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	app := gin.Default()

	// Load static resources & templates
	customTemplate := template.Must(template.New("main").ParseGlob("resources/templates/base/*.tmpl"))
	customTemplate.ParseGlob("resources/templates/*.tmpl")
	app.SetHTMLTemplate(customTemplate)
	app.Static("/static", "resources/static")

	// Homepage endpoint
	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Home",
		})
	})

	// Run on 5050 port
	app.Run(":5050")
}
