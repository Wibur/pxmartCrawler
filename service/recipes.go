package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"openCrawler/entity"
	"reflect"

	"github.com/gin-gonic/gin"
)

func FindAllRecipe(c *gin.Context) {
	recipes := entity.FindAllRecipe()
	c.JSON(http.StatusOK, recipes)
}

func FindByRecipeId(c *gin.Context) {
	recipe := entity.FindByRecipeId(c.Param("id"))
	log.Println(recipe)
	if recipe.PxId == "" {
		c.JSON(http.StatusNotFound, "Error: Recipe not found")
		return
	}
	log.Println("Recipe => ", recipe)
	c.JSON(http.StatusOK, recipe)
}

func CreateRecipes(data []byte) (res int, err error) {
	var insert map[string]interface{}
	json.Unmarshal(data, &insert)
	// code, data => list, totalCount
	fmt.Println(insert["data"])
	fmt.Println(reflect.TypeOf(insert["data"]))

	if insert["code"] != 200 {
		return 0, err
	}
	return 1, err
	// var insertData []Recipes
	// if insert["data"]["totalCount"] != 0 {
	// 	for v := range insert["data"] {

	// 	}
	// }
	// // insert data to database
	// res, err = entity.CreateRecipes()
	// return res, err
}
