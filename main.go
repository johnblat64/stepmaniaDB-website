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
	SongPage  SongPage
}

func songs(c *gin.Context) {
	response, err := http.Get("http://localhost:8080/songs")
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

	var songsPage SongPage
	err = json.Unmarshal(responseData, &songsPage)
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.ParseFiles("resources/results.gtpl"))

	data := SongsData{
		PageTitle: "Results",
		SongPage:  songsPage,
	}

	tmpl.Execute(c.Writer, data)

}

func getSong(c *gin.Context) {
	response, err := http.Get("http://localhost:8080/songs/" + c.Param("songid"))
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

	var song Song
	json.Unmarshal(responseData, &song)
	tmpl := template.Must(template.ParseFiles("resources/song.gtpl"))
	err = tmpl.Execute(c.Writer, song)
	if err != nil {
		log.Panicln(err)
	}
}

func getNotFound(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("resources/404.html"))
	tmpl.Execute(c.Writer, nil)
	c.Status(http.StatusNotFound)
}

func main() {
	r := gin.Default()
	r.Static("/resources", "./resources")
	r.GET("/", songs)
	r.GET("/songs/:songid", getSong)
	r.NoRoute(getNotFound)

	r.Run(":80")
}
