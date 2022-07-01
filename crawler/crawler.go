package crawler

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

type recipe struct {
	Keyword   string `json:"keyword"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
	Category  []int  `json:"category"`
	Type      int    `json:"type"`
}

func GetRecipes(c *gin.Context) {

	time := time.Now().UnixMilli()
	params := recipe{"", 1, 12, []int{}, 6}
	body, err := json.Marshal(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	url := getRecipeUrl + "?t=" + strconv.FormatInt(time, 10)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	jsonStr := string(body)

	c.JSON(http.StatusOK, gin.H{
		"data": jsonStr,
	})
}
