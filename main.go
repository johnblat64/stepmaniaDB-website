package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

type SongsData struct {
	PageTitle string
	Songs     []Song
}

func songs(c *gin.Context) {
	response, err := http.Get("http://localhost:8080/songList")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode >= 500 && response.StatusCode < 600 {
		c.String(500, "Internal Server Error")
		return
	}
	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var songList []Song
	err = json.Unmarshal(responseData, &songList)
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.ParseFiles("resources/results.gtpl"))

	data := SongsData{
		PageTitle: "Results",
		Songs:     songList,
	}

	tmpl.Execute(c.Writer, data)

}

func main() {
	r := gin.Default()
	r.Static("/resources", "./resources")
	r.GET("/", songs)
	r.Run(":80")
}
