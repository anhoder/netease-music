package service

import (
	"github.com/anhoder/netease-music/util"
	"net/http"
)

type VideoTimelineRecommendService struct {
	Offset string `json:"offset" form:"offset"`
}

func (service *VideoTimelineRecommendService) VideoTimelineRecommend() (float64, string) {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	if service.Offset == "" {
		data["offset"] = "0"
	} else {
		data["offset"] = service.Offset
	}
	data["filterLives"] = "[]"
	data["withProgramInfo"] = "true"
	data["needUrl"] = "1"
	data["resolution"] = "480"
	code, reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/videotimeline/get`, data, options)

	return code, reBody
}
