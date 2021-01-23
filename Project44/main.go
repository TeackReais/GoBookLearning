package main

// import (
// 	"crypto/md5"
// 	"fmt"
// 	"log"

// 	"layeh.com/radius"
// 	"layeh.com/radius/rfc2865"
// )

// func main() {
// 	handler := func(w radius.ResponseWriter, r *radius.Request) {
// 		username := rfc2865.UserName_GetString(r.Packet)
// 		password := rfc2865.UserPassword_GetString(r.Packet)

// 		var code radius.Code
// 		fmt.Println(r.Packet.Code)
// 		fmt.Println(r.Packet.Identifier)
// 		fmt.Println(r.Packet.Authenticator)
// 		fmt.Println(string(r.Packet.Secret))
// 		fmt.Println("Attributes")
// 		for _, b := range r.Packet.Attributes {
// 			fmt.Println(b.Type, b.Attribute)
// 		}
// 		fmt.Println("TruePassword", "hello")
// 		fmt.Println("TruePasswordMD5", md5.Sum([]byte("hello")))
// 		if username == "tim" && password == "12345" {
// 			code = radius.CodeAccessAccept
// 		} else {
// 			code = radius.CodeAccessReject
// 		}
// 		code = radius.CodeAccessAccept
// 		log.Printf("Writing %v to %v", code, r.RemoteAddr)
// 		w.Write(r.Response(code))
// 	}

// 	server := radius.PacketServer{
// 		Addr:         ":1812",
// 		Network:      "udp",
// 		Handler:      radius.HandlerFunc(handler),
// 		SecretSource: radius.StaticSecretSource([]byte("20010925")),
// 	}

// 	log.Printf("Starting server on :1812")
// 	if err := server.ListenAndServe(); err != nil {
// 		log.Fatal(err)
// 	}
// }
