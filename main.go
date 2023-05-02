package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
)

var GBackendApiEndpoint string

type SongsData struct {
	PageTitle        string
	SongResultsModel SongsResultsModel
}

func songs(c *gin.Context) {

	c.MultipartForm()
	queryStr := c.Request.URL.RawQuery

	response, err := http.Get(GBackendApiEndpoint + "/songs?" + queryStr)
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
	meterMinStr := c.DefaultQuery("meterMin", "0")
	songResultModel.SearchParameters.MeterMin, err = strconv.Atoi(meterMinStr)
	if err != nil {
		log.Println(err)
	}
	meterMaxStr := c.DefaultQuery("meterMax", "99")
	songResultModel.SearchParameters.MeterMax, err = strconv.Atoi(meterMaxStr)
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
	// we like to display the difficulty in order of easiest to hardest
	for _, song := range songResultModel.Songs {
		sort.Slice(song.Charts, func(i, j int) bool {
			return song.Charts[i].Meter < song.Charts[j].Meter
		})
	}
	songResultModel.PageCount = songResultsResponse.PageCount
	songResultModel.TotalSongsCount = songResultsResponse.TotalSongsCount

	tmpl := template.Must(template.ParseFiles("resources/search.gtpl", "resources/base.gtpl"))

	data := SongsData{
		PageTitle:        "Search",
		SongResultsModel: songResultModel,
	}

	tmpl.ExecuteTemplate(c.Writer, "base", data)

}

func getSong(c *gin.Context) {
	response, err := http.Get(GBackendApiEndpoint + "/songs/" + c.Param("songid"))
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
	sort.Slice(song.Charts, func(i, j int) bool {
		return song.Charts[i].Meter < song.Charts[j].Meter
	})
	tmpl := template.Must(template.ParseFiles("resources/song.gtpl", "resources/base.gtpl"))
	err = tmpl.ExecuteTemplate(c.Writer, "base", song)
	if err != nil {
		log.Panicln(err)
	}
}

func getPack(c *gin.Context) {
	response, err := http.Get(GBackendApiEndpoint + "/packs/" + c.Param("packid"))
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

	var pack Pack
	json.Unmarshal(responseData, &pack)
	tmpl := template.Must(template.ParseFiles("resources/pack.gtpl", "resources/base.gtpl"))
	err = tmpl.ExecuteTemplate(c.Writer, "base", pack)
	if err != nil {
		log.Panicln(err)
	}
}

func getNotFound(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("resources/404.html"))
	tmpl.Execute(c.Writer, nil)
	c.Status(http.StatusNotFound)
}

func getAbout(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("resources/about.gtpl", "resources/base.gtpl"))
	err := tmpl.ExecuteTemplate(c.Writer, "base", nil)
	if err != nil {
		log.Panicln(err)
	}

}

func main() {
	// get environment variable for backend api endpoint and crash if not set with log
	GBackendApiEndpoint = os.Getenv("BACKEND_API_ENDPOINT")
	if GBackendApiEndpoint == "" {
		log.Fatal("BACKEND_API_ENDPOINT environment variable not set")
	}

	r := gin.Default()
	r.Static("/resources", "./resources")
	r.GET("/", songs)
	r.GET("/songs", songs)
	r.GET("/songs/:songid", getSong)
	r.GET("/packs/:packid", getPack)
	r.GET("/about", getAbout)
	r.NoRoute(getNotFound)

	r.Run(":80")
}
