package main

import (
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

func getNotFound(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("resources/404.html"))
	tmpl.Execute(c.Writer, nil)
	c.Status(http.StatusNotFound)
}
