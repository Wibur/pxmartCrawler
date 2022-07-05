package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	getRecipeUrl = "https://www.pxmart.com.tw/Api/api/Recipes/GetHomeList"
)

type recipeParams struct {
	Keyword   string `json:"keyword" default:""`
	PageIndex int    `json:"pageIndex" default:"1"`
	PageSize  int    `json:"pageSize" default:"10"`
	Category  []int  `json:"category"`
	Type      int    `json:"type" default:"6"`
}

func CrawlRecipes(c *gin.Context) {
	time := time.Now().UnixMilli()
	var params recipeParams
	c.ShouldBind(&params)
	body, err := json.Marshal(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	url := getRecipeUrl + "?t=" + strconv.FormatInt(time, 10)

	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err = CreateRecipes(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Created Success!",
	})
}
