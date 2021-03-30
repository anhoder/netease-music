package service

import (
	"github.com/anhoder/netease-music/util"
)

type DjToplistPayService struct {
	Limit string `json:"limit" form:"limit"`
}

func (service *DjToplistPayService) DjToplistPay() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)

	if service.Limit == "" {
		data["limit"] = "10"
	} else {
		data["limit"] = service.Limit
	}

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/djradio/toplist/pay`, data, options)

	return reBody
}