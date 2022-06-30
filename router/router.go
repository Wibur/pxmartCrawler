package router

import (
	"openCrawler/service"

	"github.com/gin-gonic/gin"
)

func Register() {
	r := gin.Default()

	r.POST("/crawlRecipes", service.CrawlRecipes)
	r.GET("/getRecipe/:id", service.FindByRecipeId)
	r.GET("/getAllRecipe", service.FindAllRecipe)

	r.Run(":9527") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
