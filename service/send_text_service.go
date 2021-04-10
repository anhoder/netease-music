package service

import (
	"github.com/anhoder/netease-music/util"
	"net/http"
)

type SendTextService struct {
	ID      string `json:"playlist" form:"playlist"`
	Msg     string `json:"msg" form:"msg"`
	UserIds string `json:"user_ids" form:"user_ids"`
}

func (service *SendTextService) SendText() (float64, string) {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	data["id"] = service.ID
	data["type"] = "text"
	data["msg"] = service.Msg
	data["userIds"] = "[" + service.UserIds + "]"
	code, reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/msg/private/send`, data, options)

	return code, reBody
}
