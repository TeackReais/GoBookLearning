package main //map和json的变换

import (
	"encoding/json"
	"fmt"
)

func map2Json(map1 map[string]string) string {
	jsonStr, err := json.Marshal(map1)
	if err != nil {
		fmt.Println("Map To Json失败，请联系开发人员。")
		return ""
	}
	result := string(jsonStr)
	return result
}

func json2Map(jsonStr string) map[string]string {
	var mapResult map[string]string = make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("Map To Json失败，请联系开发人员。")
		return nil
	}
	return mapResult
}
