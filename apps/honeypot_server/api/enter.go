package api

import (
	"Honeypot/apps/honeypot_server/api/captcha_api"
	"Honeypot/apps/honeypot_server/api/user_api"
)

type Api struct {
	UserApi user_api.UserApi
	Captcha captcha_api.CaptchaApi
}

var App = new(Api)
