package main

import "net/http"

func main() {
	http.ListenAndServe("localhost:8080", nil)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
}
