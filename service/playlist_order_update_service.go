package service

import (
	"github.com/anhoder/netease-music/util"
	"net/http"
)

type PlaylistOrderUpdateService struct {
	Ids string `json:"ids" form:"ids"`
}

func (service *PlaylistOrderUpdateService) PlaylistOrderUpdate() (float64, string) {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	data["id"] = service.Ids
	code, reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/playlist/order/update`, data, options)

	return code, reBody
}
