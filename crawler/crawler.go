package crawler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	getRecipeUrl = "https://www.pxmart.com.tw/Api/api/Recipes/GetHomeList"
)

type RecipeType struct {
	Keyword   string `json:"keyword"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
	Category  []int  `json:"category"`
	Type      int    `json:"type"`
}

func GetRecipes(c *gin.Context) {
	time := time.Now().UnixMilli()
	log.Println(time)
	var ReqType RecipeType
	// params := recipe{"", 1, 12, []int{}, 6}
	c.DefaultQuery("t", string(time))
	// data, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(data)
	if err := c.ShouldBindJSON(&ReqType); err != nil {
		// c.AbortWithError(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	reqMap := map[string]interface{}{
		"code": 0,
		"rsp":  fmt.Sprintf("welcome %v", ReqType),
	}

	c.JSON(http.StatusOK, reqMap)
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": msg,
	})
}
