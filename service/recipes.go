package service

import (
	"encoding/json"
	"errors"
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

func CreateRecipes(data []byte) error {
	var insert entity.ResponseBody
	if err := json.Unmarshal(data, &insert); err != nil {
		return err
	}

	if insert.Code != http.StatusOK {
		return errors.New("Third Party Error")
	}

	if err := entity.CreateRecipes(insert.Data.List); err != nil {
		return err
	}

	return nil
}
