package api

import (
	"honeypot_server/internal/api/captcha_api"
	"honeypot_server/internal/api/log_api"
	"honeypot_server/internal/api/user_api"
)

type Api struct {
	UserApi    user_api.UserApi
	CaptchaApi captcha_api.CaptchaApi
	LogApi     log_api.LogApi
}

var App = new(Api)
