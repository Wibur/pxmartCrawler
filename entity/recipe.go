package entity

import (
	"openCrawler/databases"
)

/*
map[
	code:200
	data:map[
		list:[
			map[cookingTime:15 id:35e0e2d0-743a-4a92-81f8-29ff120e7777 image:https://www.pxmart.com.tw/Api/Images/132966373658821913.JPG ingredientNum:5 name:菲力牛柳佐椰香綠咖哩 ]
			map[cookingTime:20 id:a98373b0-b43d-46ae-9bc9-030ce4b7c69a image:https://www.pxmart.com.tw/Api/Images/132966370711244501.JPG ingredientNum:5 name:黃咖哩大蝦 ]
			map[cookingTime:20 id:b1ac73c1-1a53-4631-aba7-c535d5012a0b image:https://www.pxmart.com.tw/Api/Images/132966367140627027.JPG ingredientNum:8 name:爪哇雞腿排咖哩    ]
			map[cookingTime:<nil> id:5fcc7e18-0e28-446a-a597-b7d0436625d5 image:https://www.pxmart.com.tw/Api/Images/132953295664145280.JPG ingredientNum:2 name:親子手作便當]
			map[cookingTime:30 id:562fe6a4-cb4f-496c-8031-faa2e3ba6549 image:https://www.pxmart.com.tw/Api/Images/132953291489243580.JPG ingredientNum:6 name:韓式豬肉大醬湯]
			map[cookingTime:20 id:fe752852-5945-4258-85fa-24a6152463b0 image:https://www.pxmart.com.tw/Api/Images/132953289432138522.jpg ingredientNum:1 name:韓式辣味豆芽菜]
			map[cookingTime:20 id:05dc39da-8f9d-47c8-8d26-bf5357a996d2 image:https://www.pxmart.com.tw/Api/Images/132953286692786945.jpg ingredientNum:1 name:韓式涼拌醃黃瓜]
			map[cookingTime:20 id:a0e89be9-7ecc-473e-b8e8-6e5353461d8d image:https://www.pxmart.com.tw/Api/Images/132953283653807061.jpg ingredientNum:1 name:韓式梅汁小蕃茄]
			map[cookingTime:60 id:78c3d9f8-2aee-4a0a-8f58-cc319cd222bf image:https://www.pxmart.com.tw/Api/Images/132857357237573149.JPG ingredientNum:8 name:蕈菇咖哩南瓜盅筆管麵]
			map[cookingTime:15 id:e1b2d391-97c9-425a-a6bd-8d374db0eee0 image:https://www.pxmart.com.tw/Api/Images/132853222361060829.jpg ingredientNum:6 name:東北酸白御膳鍋]
			map[cookingTime:20 id:5a5b3336-3f00-4f17-bb09-753207fc974e image:https://www.pxmart.com.tw/Api/Images/132853218442056660.jpg ingredientNum:6 name:澱粉炸彈韓式部隊拉麵]
			map[cookingTime:40 id:36f6184e-f591-4ff1-bb4f-8684ae99ed70 image:https://www.pxmart.com.tw/Api/Images/132774544128796803.jpg ingredientNum:7 name:氣炸鯖魚蔬菜好食]
		]
		totalCount:28
	]
	message:
]
*/

type ResponseBody struct {
	Code    int `json:"code"`
	Data    ResponseData
	Message string `json:"message"`
}

type ResponseData struct {
	List       []Recipes
	TotalCount int `json:"totalCount"`
}

type Recipes struct {
	PxId        string  `json:"id"`
	CookingTime float64 `json:"cookingTime"`
	Image       string  `json:"image"`
	Ingredient  int     `json:"ingredientNum"`
	Name        string  `json:"name"`
}

func FindAllRecipe() []Recipes {
	var recipes []Recipes
	databases.DBConnect.Find(&recipes)
	return recipes
}

func FindByRecipeId(id string) Recipes {
	var recipe Recipes
	databases.DBConnect.Where("px_id = ?", id).First(&recipe)
	return recipe
}

func CreateRecipes(recipes []Recipes) {
	databases.DBConnect.Create(&recipes)
}
