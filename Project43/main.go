package main

import (
	"fmt"
	"io/ioutil"
	// "net/http"
	// "strconv"
	// "time"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	fmt.Println("QRCODE")
	var png []byte
	png, err := qrcode.Encode("https://www.baidu.com", qrcode.Highest, 256)
	if err != nil {
		fmt.Println("ERR=", err)
	}
	err = ioutil.WriteFile("QRCODE.png", png, 0600)
	if err != nil {
		fmt.Println("ERR=", err)
	}
	// http.HandleFunc("/login/", loginHandle)
	// http.HandleFunc("/start/", startHandle)
	// server := &http.Server{
	// 	Addr:    "localhost.treais.net:8080",
	// 	Handler: nil,
	// }
	// server.ListenAndServe()
}

// func startHandle(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	LastingTime, err := strconv.Atoi(r.Form.Get("time"))
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	go ShowQRCODE(w, r)
// 	time.Sleep(time.Duration(LastingTime) * time.Second)
// 	return
// }

// func loginHandle(w http.ResponseWriter, r *http.Request) {

// }

// func ShowQRCODE(w http.ResponseWriter, r *http.Request) {

// }
