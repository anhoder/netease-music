package service

import (
	"github.com/anhoder/netease-music/util"
	"net/http"
)

type PersonalizedService struct {
	Limit string `json:"limit" form:"limit"`
}

func (service *PersonalizedService) Personalized() (float64, string) {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	if service.Limit == "" {
		data["limit"] = "30"
	} else {
		data["limit"] = service.Limit
	}
	data["order"] = "true"
	data["n"] = "1000"
	code, reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/personalized/playlist`, data, options)

	return code, reBody
}
