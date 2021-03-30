package service

import (
	"github.com/anhoder/netease-music/util"
)

type ArtistsService struct {
	ID string `json:"id" form:"id"`
}

func (service *ArtistsService) Artists() map[string]interface{} {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["id"] = service.ID

	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/v1/artist/`+service.ID, data, options)

	return reBody
}