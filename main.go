package main

import (
	"weeb-dex-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/manga", controllers.GetManga)

	r.Run()
}
