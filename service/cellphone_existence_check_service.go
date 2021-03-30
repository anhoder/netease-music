package service

import (
	"github.com/anhoder/netease-music/util"
)

type CellphoneExistenceCheckService struct {
	Cellphone   string `json:"phone" form:"phone"`
	Countrycode string `json:"countrycode" form:"countrycode"`
}

func (service *CellphoneExistenceCheckService) CellphoneExistenceCheck() map[string]interface{} {

	options := &util.Options{
		Crypto:  "eapi",
		Url:     "/api/cellphone/existence/check",
	}
	data := make(map[string]string)
	if service.Countrycode != "" {
		data["countrycode"] = service.Countrycode
	}
	data["cellphone"] = service.Cellphone

	reBody, _ := util.CreateRequest("POST", `http://music.163.com/eapi/cellphone/existence/check`, data, options)

	return reBody
}