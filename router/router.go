package router

import (
	"openCrawler/crawler"
	"openCrawler/service"

	"github.com/gin-gonic/gin"
)

func Register() {
	r := gin.Default()

	r.POST("/createRecipes", crawler.GetRecipes)
	r.GET("/getRecipe/:id", service.FindByRecipeId)
	r.GET("/getAllRecipe", service.FindAllRecipe)

	r.Run(":9527") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
