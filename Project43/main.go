package main

import (
	"fmt"
	"net/http"
	"os"
	"path"

	// "io/ioutil"
	// "net/http"
	// "strconv"
	// "time"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	var information string
	ProgramPath, err := os.Getwd()
	Path := path.Dir(ProgramPath) + "/File.png"
	fmt.Println("请输入二维码的内容，程序将生成一张二维码图片")
	fmt.Scanln(&information)
	fmt.Println(Path)
	data, err := qrcode.Encode(information, qrcode.Highest, 512)
	if err != nil {
		fmt.Println("ERR=", err)
		return
	}
	fmt.Println("Success")
	http.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(data)
	})
	http.ListenAndServe("localhost:8080", nil)
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
