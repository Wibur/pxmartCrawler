package router

import (
	"openCrawler/crawler"

	"github.com/gin-gonic/gin"
)

func Register() {
	r := gin.Default()
	
	r.POST("/getRecipes", crawler.GetRecipes)

	r.Run(":9527") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
