package router

import (
	crawl "openCrawler/crawler"

	"github.com/gin-gonic/gin"
)

func Register() {
	r := gin.Default()
	// crawler
	// r.POST("/getAssets", crawl.GetAssets)
	r.POST("/getRecipes", crawl.GetRecipes)
	// r.POST("/getAssets", func(c *gin.Context) {
	// 	// log.Println(c.PostForm("url"))
	// 	params := crawlerParams{c.DefaultPostForm("url", "")}
	// 	// log.Println(reflect.TypeOf(params))
	// 	log.Println(params)
	// 	// go to crawl
	// 	crawl.Fetch(params.url)
	// })
	r.Run(":9527") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
