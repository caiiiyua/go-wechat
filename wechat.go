package main

import (
	"log"
	"net/http"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/message/request"
	"github.com/chanxuehong/wechat/util"
)

const token string = "nehe"
const appid string = "wxe039839fbc011f6d"
const appsecret string = "929194a70c73eed9f8bec14528a2b8c2"
const wechatid string = "gh_059a1b6286af"
const aesKeyOrigin = "KI5r9bVLmV5JiWiVlLiAUFpvCZHEG0wxxEp2lnzNeQT"

func textMessageHandler(w http.ResponseWriter, r *mp.Request) {

}

func invalideRequestHandler(w http.ResponseWriter, r *http.Request, err error) {

}

func main() {
	aesKey, err := util.AESKeyDecode(aesKeyOrigin)
	if err != nil {
		log.Fatalln(err)
	}

	messageServeMux := mp.NewMessageServeMux()
	messageServeMux.MessageHandleFunc(request.MsgTypeText, textMessageHandler)

	wechatServer := mp.NewDefaultWechatServer(wechatid, token, appid, aesKey,
		messageServeMux)
	wechatServerFrontend := mp.NewWechatServerFrontend(wechatServer,
		mp.InvalidRequestHandlerFunc(invalideRequestHandler))

	http.Handle("/wechat", wechatServerFrontend)
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Start wechat server failed")
	} else {
		log.Println("Start wechat server and listening on 80 ...")
	}
}
