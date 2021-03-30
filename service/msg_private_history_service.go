package service

import (
	"net/http"
	"github.com/anhoder/netease-music/util"
)

type MsgPrivateHistoryService struct {
	UID   string `json:"uid" form:"uid"`
	Limit string `json:"limit" form:"limit"`
	Time  string `json:"before" form:"before"`
}

func (service *MsgPrivateHistoryService) MsgPrivateHistory() map[string]interface{} {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	data["userId"] = service.UID
	if service.Limit == "" {
		data["limit"] = "30"
	} else {
		data["limit"] = service.Limit
	}
	if service.Time == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Time
	}
	data["order"] = "true"
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/msg/private/history`, data, options)

	return reBody
}