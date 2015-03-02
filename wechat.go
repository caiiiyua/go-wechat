package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/message/request"
	"github.com/chanxuehong/wechat/mp/message/response"
	"github.com/chanxuehong/wechat/util"
)

const token string = "nehe"
const appid string = "wxe039839fbc011f6d"
const appsecret string = "929194a70c73eed9f8bec14528a2b8c2"
const wechatid string = "gh_059a1b6286af"
const aesKeyOrigin = "KI5r9bVLmV5JiWiVlLiAUFpvCZHEG0wxxEp2lnzNeQT"

func textMessageHandler(w http.ResponseWriter, r *mp.Request) {
	text := request.GetText(r.MixedMsg)
	// resp := response.NewText(text.FromUserName, text.ToUserName, text.CreateTime,
	// text.Content)
	articles := make([]response.Article, 3)
	about := response.Article{
		"我的博客",
		"点击阅读博客文章",
		"https://mmbiz.qlogo.cn/mmbiz/kqQvK5zClJIT4InuQGaNzJHgC0Al4Kib3wNicXq1HFTboeZO0HiagteQGCdgbVM4Rnr8NKMiaricRSN8zMGRT1z91fg/0",
		"http://www.imogu.us/"}
	articles = append(articles, about)
	resp := response.NewNews(text.FromUserName, text.ToUserName, text.CreateTime,
		articles)
	mp.WriteRawResponse(w, r, resp)
}

func invalideRequestHandler(w http.ResponseWriter, r *http.Request, err error) {
	io.WriteString(w, err.Error())
	fmt.Println("invalide request")
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
