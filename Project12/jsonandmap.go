package main

import (
	"encoding/json"
)

func map2Json(map1 map[string]string) string {
	jsonStr, err := json.Marshal(map1)
	errorExit(err, 5)
	return string(jsonStr)
}

func json2Map(jsonStr string) map[string]string {
	var mapResult map[string]string = make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	errorExit(err, 6)
	return mapResult
}
