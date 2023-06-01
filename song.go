package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
	smdbcore "github.com/stepmaniadb/stepmaniadb-core"
)

func generateBannerUrl(bannerPath string) string {
	if bannerPath == "" {
		return ""
	}
	return GEnvironmentConfig.FileStoreUrl + "/" + bannerPath
}

type SongResultsResponse struct {
	Page            int             `json:"pageNum"`
	PageSize        int             `json:"pageSize"`
	PageCount       int             `json:"pageCount"`
	TotalSongsCount int             `json:"totalSongsCount" db:"total_songs_count"`
	Songs           []smdbcore.Song `json:"songs"`
}

type SongSearchParameters struct {
	Title                    string
	Artist                   string
	Credit                   string
	Pack                     string
	TimeSignatureNumerator   int
	TimeSignatureDenominator int
	BpmMin                   int
	BpmMax                   int
	MeterMin                 int
	MeterMax                 int
	StepsType                string
}

// for HTML page
type SongsResultsModel struct {
	FileStoreUrl     string
	Page             int             `json:"pageNum"`
	PageSize         int             `json:"pageSize"`
	PageCount        int             `json:"pageCount"`
	TotalSongsCount  int             `json:"totalSongsCount"`
	Songs            []smdbcore.Song `json:"songs"`
	StepsTypeOptions []string
	SearchParameters SongSearchParameters
}

type SongModel struct {
	Song smdbcore.Song `json:"song"`
}

func (params SongSearchParameters) AsQueryString() string {
	return "?title=" + params.Title + "&artist=" + params.Artist + "&credit=" + params.Credit + "&pack=" + params.Pack + "&stepstype=" + params.StepsType + "&timeSignatureNumerator=" + strconv.Itoa(params.TimeSignatureNumerator) + "&timeSignatureDenominator=" + strconv.Itoa(params.TimeSignatureDenominator) + "&bpmMin=" + strconv.Itoa(params.BpmMin) + "&bpmMax=" + strconv.Itoa(params.BpmMax) + "&meterMin=" + strconv.Itoa(params.MeterMin) + "&meterMax=" + strconv.Itoa(params.MeterMax)
}

func (resultsModel SongsResultsModel) NextPage() int {
	return resultsModel.Page + 1
}

func (resultsModel SongsResultsModel) HasNextPage() bool {
	return resultsModel.Page < resultsModel.PageCount
}

func (resultsModel SongsResultsModel) PreviousPage() int {
	return resultsModel.Page - 1
}

// struct for counting songs
type Count struct {
	Count int `db:"count"`
}

type SongsData struct {
	PageTitle        string
	SongResultsModel SongsResultsModel
}

// Endpoints

func getSongs(c *gin.Context) {

	c.MultipartForm()
	queryStr := c.Request.URL.RawQuery

	response, err := http.Get(GEnvironmentConfig.BackendApiUrl + "/songs?" + queryStr)
	if err != nil {
		log.Println(err)
		c.String(500, "An Error Occurred")
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 500 && response.StatusCode < 600 {
		log.Println("Server Error: " + response.Status)
		c.String(500, "Internal Server Error")
		return
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		c.String(500, "An Error Occurred")
		return
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
		log.Println(err)
		c.String(500, "An Error Occurred")
		return
	}

	songResultModel.FileStoreUrl = GEnvironmentConfig.FileStoreUrl

	songResultModel.Songs = songResultsResponse.Songs
	// we like to display the difficulty in order of easiest to hardest
	for _, song := range songResultModel.Songs {
		sort.Slice(song.Charts, func(i, j int) bool {
			return song.Charts[i].Meter < song.Charts[j].Meter
		})
	}
	songResultModel.PageCount = songResultsResponse.PageCount
	songResultModel.TotalSongsCount = songResultsResponse.TotalSongsCount

	tmpl := template.Must(template.New("resources/search.gtpl").Funcs(template.FuncMap{
		"generateBannerUrl": generateBannerUrl,
	}).ParseFiles("resources/search.gtpl", "resources/base.gtpl"))

	data := SongsData{
		PageTitle:        "Search",
		SongResultsModel: songResultModel,
	}

	err = tmpl.ExecuteTemplate(c.Writer, "base", data)

	if err != nil {
		log.Println(err)
		c.String(500, "An Error Occurred")
		return
	}

}

func getSong(c *gin.Context) {
	response, err := http.Get(GEnvironmentConfig.BackendApiUrl + "/songs/" + c.Param("songid"))
	if err != nil {
		log.Println(err)
		c.String(500, "An Error Occurred")
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 500 && response.StatusCode < 600 {
		c.String(500, "Internal Server Error")
		return
	}
	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Println(err)
		c.String(500, "An Error Occurred")
		return
	}

	var song smdbcore.Song
	json.Unmarshal(responseData, &song)
	sort.Slice(song.Charts, func(i, j int) bool {
		return song.Charts[i].Meter < song.Charts[j].Meter
	})

	tmpl := template.Must(template.New("resources/song.gtpl").Funcs(template.FuncMap{
		"generateBannerUrl": generateBannerUrl,
	}).ParseFiles("resources/song.gtpl", "resources/base.gtpl"))

	err = tmpl.ExecuteTemplate(c.Writer, "base", song)
	if err != nil {
		log.Println(err)
		c.String(500, "An Error Occurred")
		return
	}
}
