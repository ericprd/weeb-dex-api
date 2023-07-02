package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetManga(c *gin.Context) {
	// var manga []models.MangaDataStruct
	err := godotenv.Load()
	var BASE_URL = os.Getenv("BASE_URL")
	data, err := http.Get(BASE_URL+"manga")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error": err.Error() })
		return
	}

	defer data.Body.Close()
	body, err := ioutil.ReadAll(data.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error": err.Error() })
		return
	}

	var bodyData = map[string]any{}
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error": err.Error() })
		return
	}
	refinedData, ok := bodyData["data"].([]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H { "error": "invalid" })
		return
	}
	c.JSON(http.StatusOK, gin.H { "data": refinedData })
}
