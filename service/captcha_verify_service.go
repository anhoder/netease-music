package service

import (
	"github.com/anhoder/netease-music/util"
	"net/http"
)

type CaptchaVerifyService struct {
	Ctcode    string `json:"ctcode" form:"ctcode"`
	Cellphone string `json:"phone" form:"phone"`
	Captcha   string `json:"captcha" form:"captcha"`
}

func (service *CaptchaVerifyService) CaptchaVerify() (float64, string) {

	cookiesOS := &http.Cookie{Name: "os", Value: "pc"}

	options := &util.Options{
		Crypto:  "weapi",
		Cookies: []*http.Cookie{cookiesOS},
	}
	data := make(map[string]string)
	if service.Ctcode == "" {
		data["ctcode"] = "86"
	} else {
		data["ctcode"] = service.Ctcode
	}
	data["cellphone"] = service.Cellphone
	data["captcha"] = service.Captcha

	code, reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/sms/captcha/verify`, data, options)

	return code, reBody
}
