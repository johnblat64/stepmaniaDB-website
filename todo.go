package main

import (
	"log"
	"text/template"

	"github.com/gin-gonic/gin"
)

func getTodo(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("resources/todo.gtpl", "resources/base.gtpl"))
	err := tmpl.ExecuteTemplate(c.Writer, "base", nil)
	if err != nil {
		log.Panicln(err)
	}

}
