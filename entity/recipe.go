package entity

import (
	"openCrawler/databases"
)

type Recipes struct {
	Id          int     `json:"Id"`
	PxId        string  `json:"PxId"`
	CookingTime float64 `json:"CookingTime"`
	Image       string  `json:"Image"`
	Ingredient  int     `json:"IngredientNum"`
	Name        string  `json:"Name"`
}

func FindAllRecipe() []Recipes {
	var recipes []Recipes
	databases.DBConnect.Find(&recipes)
	return recipes
}

func FindByRecipeId(id string) Recipes {
	var recipe Recipes
	databases.DBConnect.Where("id = ?", id).First(&recipe)
	return recipe
}

func CreateRecipes(recipes []Recipes) {

}
