package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"weeb-dex-api/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type responseStruct struct {
	Data []models.MangaDataStruct `json:"data"`
}

var _ = godotenv.Load()
var BASE_URL = os.Getenv("BASE_URL")

func GetManga(c *gin.Context) {
	var response responseStruct
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

	json.Unmarshal(body, &response)

	c.JSON(http.StatusOK, gin.H { "data": response.Data })
}
