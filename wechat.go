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
		log.Fatalln("Reuqest is not a valid request from Wechat!")
		return false
	}
	fmt.Printf("hello wechat! request[%s, %s, %s, %s]", signature, timestamp, nonce,
		echostr)
	fmt.Fprintf(w, echostr)
	log.Println("Validate request from Wechat successfully!")
	return true
}

func wechatHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	echostr := r.Form.Get("echostr")
	if len(echostr) > 0 {
		if !validate(w, r) {
			log.Fatalln("validate failed")
		}
	}
	fmt.Fprintf(w, "default handler")
	fmt.Println("default handler")
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
