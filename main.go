package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
)

type SongsData struct {
	PageTitle        string
	SongResultsModel SongsResultsModel
}

func songs(c *gin.Context) {

	c.MultipartForm()
	queryStr := c.Request.URL.RawQuery

	response, err := http.Get("http://localhost:8080/songs?" + queryStr)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode >= 500 && response.StatusCode < 600 {
		log.Println("Server Error: " + response.Status)
		c.String(500, "Internal Server Error")
		return
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var songResultModel SongsResultsModel
	songResultModel.StepsTypeOptions = []string{"dance-single", "dance-double", "dance-couple", "dance-solo", "dance-threepanel", "dance-routine", "pump-single", "pump-halfdouble", "pump-double", "pump-couple", "pump-routine", "kb7-single", "kb7-small", "ez2-single", "ez2-double", "ez2-real", "para-single", "para-double", "beat-single5", "beat-single7", "beat-double5", "beat-double7", "techno-single4", "techno-single5", "techno-single8", "techno-double4", "techno-double5", "techno-double8", "pnm-five", "pnm-nine", "kickbox-human", "kickbox-quadarm", "kickbox-insect", "kickbox-arachnid"}
	songResultModel.SearchParameters.Title = c.Query("title")
	songResultModel.SearchParameters.Artist = c.Query("artist")
	songResultModel.SearchParameters.Credit = c.Query("credit")
	songResultModel.SearchParameters.Pack = c.Query("pack")
	songResultModel.SearchParameters.StepsType = c.Query("stepstype")
	timeSignatureNumeratorStr := c.DefaultQuery("timeSignatureNumerator", "4")
	songResultModel.SearchParameters.TimeSignatureNumerator, err = strconv.Atoi(timeSignatureNumeratorStr)
	if err != nil {
		log.Println(err)
	}
	timeSignatureDenominatorStr := c.DefaultQuery("timeSignatureDenominator", "4")
	songResultModel.SearchParameters.TimeSignatureDenominator, err = strconv.Atoi(timeSignatureDenominatorStr)
	if err != nil {
		log.Println(err)
	}
	bpmMinStr := c.DefaultQuery("bpmMin", "0")
	songResultModel.SearchParameters.BpmMin, err = strconv.Atoi(bpmMinStr)
	if err != nil {
		log.Println(err)
	}
	bpmMaxStr := c.DefaultQuery("bpmMax", "999")
	songResultModel.SearchParameters.BpmMax, err = strconv.Atoi(bpmMaxStr)
	if err != nil {
		log.Println(err)
	}
	pageStr := c.DefaultQuery("page", "1")
	songResultModel.Page, err = strconv.Atoi(pageStr)
	if err != nil {
		log.Println(err)
	}
	pageSizeStr := c.DefaultQuery("pageSize", "20")
	songResultModel.PageSize, err = strconv.Atoi(pageSizeStr)
	if err != nil {
		log.Println(err)
	}

	var songResultsResponse SongResultsResponse
	err = json.Unmarshal(responseData, &songResultsResponse)
	if err != nil {
		log.Fatal(err)
	}

	songResultModel.Songs = songResultsResponse.Songs
	songResultModel.PageCount = songResultsResponse.PageCount
	songResultModel.TotalSongsCount = songResultsResponse.TotalSongsCount

	tmpl := template.Must(template.ParseFiles("resources/results.gtpl"))

	data := SongsData{
		PageTitle:        "Results",
		SongResultsModel: songResultModel,
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
	r.GET("/songs", songs)
	r.GET("/songs/:songid", getSong)
	r.NoRoute(getNotFound)

	r.Run(":80")
}
