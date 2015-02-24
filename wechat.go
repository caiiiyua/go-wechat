package main

import (
	"fmt"
	"log"
	"net/http"
)

func wechatHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	signature := r.Form.Get("signature")
	timestamp := r.Form.Get("timestamp")
	nonce := r.Form.Get("nonce")
	echostr := r.Form.Get("echostr")
	fmt.Printf("hello wechat! request[%s, %s, %s, %s]", signature, timestamp, nonce,
		echostr)
}

func main() {
	http.HandleFunc("/wechat", wechatHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Start wechat server failed")
	} else {
		log.Println("Start wechat server and listening on 80 ...")
	}
}
