package service

import (
	"github.com/anhoder/netease-music/util"
)

type VideoDetailService struct {
	ID string `json:"id" form:"id"`
}

func (service *VideoDetailService) VideoDetail() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["id"] = service.ID
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/cloudvideo/v1/video/detail`, data, options)

	return reBody
}