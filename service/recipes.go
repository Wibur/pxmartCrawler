package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"openCrawler/entity"

	"github.com/gin-gonic/gin"
)

func FindAllRecipe(c *gin.Context) {
	recipes := entity.FindAllRecipe()
	c.JSON(http.StatusOK, gin.H{
		"data":    recipes,
		"message": "Success",
	})
}

func FindByRecipeId(c *gin.Context) {
	recipe := entity.FindByRecipeId(c.Param("id"))
	err := errors.New("Recipe not found")
	if recipe.PxId == "" {
		err = fmt.Errorf("Service FindByRecipeId => %w", err)
		c.JSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    recipe,
		"message": "Success",
	})
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
