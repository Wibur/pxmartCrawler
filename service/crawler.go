package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	getRecipeUrl = "https://www.pxmart.com.tw/Api/api/Recipes/GetHomeList"
)

type recipeParams struct {
	Keyword   string `json:"keyword"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
	Category  []int  `json:"category"`
	Type      int    `json:"type"`
}

func CrawlRecipes(c *gin.Context) {
	time := time.Now().UnixMilli()
	params := recipeParams{"", 1, 12, []int{}, 6}
	body, err := json.Marshal(params)
	if err != nil {
		log.Println(err.Error())
	}
	url := getRecipeUrl + "?t=" + string(time)

	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Println(err.Error())
	}

	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
	}

	if err = CreateRecipes(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Created Success!",
	})
}
