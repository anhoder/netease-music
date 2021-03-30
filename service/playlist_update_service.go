package service

import (
	"net/http"
	"github.com/anhoder/netease-music/util"
)

type PlaylistUpdateService struct {
	Id   string `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Desc string `json:"desc" form:"desc"`
	Tags string `json:"tags" form:"tags"`
}

func (service *PlaylistUpdateService) PlaylistUpdate() map[string]interface{} {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	data["/api/playlist/desc/update"] = `{"id":` + service.Id + `,"desc":"` + service.Desc + `"}`
	data["/api/playlist/tags/update"] = `{"id":` + service.Id + `,"tags":"` + service.Tags + `"}`
	data["/api/playlist/update/name"] = `{"id":` + service.Id + `,"name":"` + service.Name + `"}`
	reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/batch`, data, options)

	return reBody
}