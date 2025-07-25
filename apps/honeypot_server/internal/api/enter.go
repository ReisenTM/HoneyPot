package api

import (
	"Honeypot/apps/honeypot_server/api/captcha_api"
	"Honeypot/apps/honeypot_server/api/log_api"
	"Honeypot/apps/honeypot_server/api/user_api"
)

type Api struct {
	UserApi    user_api.UserApi
	CaptchaApi captcha_api.CaptchaApi
	LogApi     log_api.LogApi
}

var App = new(Api)
