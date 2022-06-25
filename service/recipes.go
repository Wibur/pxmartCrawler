package service

import (
	"encoding/json"
	"fmt"
)

func CreateRecipes(data []byte) {
	var result map[string]interface{}
	json.Unmarshal(data, &result)
	// code, data => list, totalCount
	fmt.Println(result["data"], result["code"])

	// insert data to database
}
