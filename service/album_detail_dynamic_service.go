package service

import (
	"github.com/anhoder/netease-music/util"
)

type AlbumDetailDynamicService struct {
	ID string `json:"id" form:"id"`
}

func (service *AlbumDetailDynamicService) AlbumDetailDynamic() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["id"] = service.ID
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/album/detail/dynamic`, data, options)

	return reBody
}