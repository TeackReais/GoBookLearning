package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type User struct {
	ID       int
	Account  string
	Password string
}

func main() {
	user1 := User{
		ID:       1,
		Account:  "Hello",
		Password: "World",
	}
	data, _ := xml.Marshal(user1)
	fmt.Println(string(data))
	http.HandleFunc("/api/", handleRequest)
	data = []byte(xml.Header + string(data))
	err := http.ListenAndServe("localhost.treais.net:8060", nil)
	fmt.Println(err)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(getROOTRUL(r.URL.Path, 2)))
	fmt.Println(getROOTRUL(r.URL.Path, 2))
}

func getROOTRUL(url string, n int) string {
	var number int
	var runes []rune = make([]rune, 1)
	for _, runeEle := range url {
		if runeEle == '/' {
			number++
			if number == n {
				return string(runes)
			}
		} else {
			runes = append(runes, runeEle)
		}
	}
	if number == n-1 {
		return string(runes)
	}
	return ""
}
