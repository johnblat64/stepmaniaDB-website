package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	smdbcore "github.com/stepmaniadb/stepmaniadb-core"
)

func getPack(c *gin.Context) {
	response, err := http.Get(GEnvironmentConfig.BackendApiUrl + "/packs/" + c.Param("packid"))
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

	var pack smdbcore.Pack
	json.Unmarshal(responseData, &pack)

	tmpl := template.Must(template.New("resources/pack.gtpl").Funcs(template.FuncMap{
		"generateBannerUrl": generateBannerUrl,
	}).ParseFiles("resources/pack.gtpl", "resources/base.gtpl"))

	err = tmpl.ExecuteTemplate(c.Writer, "base", pack)
	if err != nil {
		log.Println(err)
		c.String(500, "An Error Occurred")
		return
	}
}
