package service

import (
	"github.com/anhoder/netease-music/util"
)

type PersonalizedDjprogramService struct {
	ID     string `json:"id" form:"id"`
	Limit  string `json:"limit" form:"limit"`
	Offset string `json:"offset" form:"offset"`
}

func (service *PersonalizedDjprogramService) PersonalizedDjprogram() (float64, string) {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)

	code, reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/personalized/djprogram`, data, options)

	return code, reBody
}
