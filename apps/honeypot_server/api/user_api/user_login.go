package user_api

import (
	"Honeypot/apps/honeypot_server/global"
	"Honeypot/apps/honeypot_server/middleware"
	"Honeypot/apps/honeypot_server/models"
	"Honeypot/apps/honeypot_server/utils/captcha"
	"Honeypot/apps/honeypot_server/utils/jwts"
	"Honeypot/apps/honeypot_server/utils/pwd"
	"Honeypot/apps/honeypot_server/utils/resp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type UserLoginRequest struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required" label:"密码"`
	CaptchaID   string `json:"captcha_id" binding:"required"`
	CaptchaCode string `json:"captcha_code" binding:"required"`
}

func (UserApi) UserLoginView(c *gin.Context) {
	cr := middleware.GetBind[UserLoginRequest](c)
	if cr.CaptchaID == "" || cr.CaptchaCode == "" {
		resp.FailWithMsg("请输入图片验证码", c)
		return
	}
	if !captcha.CaptchaStore.Verify(cr.CaptchaID, cr.CaptchaCode, true) {
		resp.FailWithMsg("图片验证码验证失败", c)
		return
	}
	//颁发token
	var user models.UserModel
	err := global.DB.Take(&user, "username = ?", cr.Username).Error
	if err != nil {
		resp.FailWithMsg("用户名或密码错误", c)
		return
	}

	if !pwd.CompareHashAndPassword(user.Password, cr.Password) {
		resp.FailWithMsg("用户名或密码错误", c)
		return
	}

	token, err := jwts.GetToken(jwts.Claims{
		UserID: user.ID,
		Role:   user.Role,
	})
	if err != nil {
		logrus.Errorf("生成token失败 %s", err)
		resp.FailWithMsg("登录失败", c)
		return
	}

	now := time.Now().Format(time.DateTime)
	global.DB.Model(&user).Update("last_login_date", now)

	resp.OkWithData(token, c)
}
