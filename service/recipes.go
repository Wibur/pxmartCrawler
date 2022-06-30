package service

import (
	"encoding/json"
	"net/http"
	"openCrawler/entity"

	"github.com/gin-gonic/gin"
)

func FindAllRecipe(c *gin.Context) {
	recipes := entity.FindAllRecipe()
	c.JSON(http.StatusOK, recipes)
}

func FindByRecipeId(c *gin.Context) {
	recipe := entity.FindByRecipeId(c.Param("id"))

	if recipe.PxId == "" {
		c.JSON(http.StatusNotFound, "Error: Recipe not found")
		return
	}

	c.JSON(http.StatusOK, recipe)
}

func CreateRecipes(data []byte) bool {
	var insert entity.ResponseBody
	json.Unmarshal(data, &insert)

	if insert.Code != http.StatusOK {
		return false
	}

	return entity.CreateRecipes(insert.Data.List)
}
