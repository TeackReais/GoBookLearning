package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe("localhost:8080", nil)
	for {
		var input string
		fmt.Scanln(&input)
		http.HandleFunc("/"+input, GenerateHandler(input))
	}
}

func GenerateHandler(pattern string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(pattern))
	}
}
