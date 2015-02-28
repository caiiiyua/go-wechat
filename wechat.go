package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
)

const token string = "nehe"
const appid string = "wxe039839fbc011f6d"
const appsecret string = "929194a70c73eed9f8bec14528a2b8c2"

func makeSignature(t string, ts string, nonce string) string {
	s1 := []string{t, ts, nonce}
	sort.Strings(s1)
	h := sha1.New()
	io.WriteString(h, strings.Join(s1, ""))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func validate(w http.ResponseWriter, r *http.Request) bool {
	signature := r.Form.Get("signature")
	timestamp := r.Form.Get("timestamp")
	nonce := r.Form.Get("nonce")
	echostr := r.Form.Get("echostr")
	if signature != makeSignature(token, timestamp, nonce) {
		log.Println("Reuqest is not a valid request from Wechat!")
		return false
	}
	if len(echostr) > 0 {
		fmt.Fprintf(w, echostr)
	}
	log.Println("Validate request from Wechat successfully!")
	return true
}

func respBuilder(category string) string {
	return "test"
}

// postMsgHandler is for "POST" request of http
func postMsgHandler(w http.ResponseWriter, r *http.Request) {

}

// messageHandler is handling all of messages from wechat
func messageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Println("POST request", r)
		postMsgHandler(w, r)
	default:
		fmt.Println("default message handler")
	}

}

func wechatHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if !validate(w, r) {
		fmt.Fprintf(w, "404 not found message!")
		return
	}
	fmt.Fprintf(w, "default handler")
	fmt.Println("default handler")

	// message handler
	messageHandler(w, r)
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
