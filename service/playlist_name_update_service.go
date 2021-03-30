package service

import (
	"github.com/anhoder/netease-music/util"
)

type PlaylistNameUpdateService struct {
	Id   string `json:"id" form:"id"`
	Name string `json:"desc" form:"name"`
}

func (service *PlaylistNameUpdateService) PlaylistNameUpdate() map[string]interface{} {

	options := &util.Options{
		Crypto:  "eapi",
		Url:     "/api/playlist/update/name",
	}
	data := make(map[string]string)
	data["id"] = service.Id
	data["name"] = service.Name
	reBody, _ := util.CreateRequest("POST", `http://interface3.music.163.com/eapi/playlist/update/name`, data, options)

	return reBody
}