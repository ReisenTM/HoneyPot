package user_api

import (
	"Honeypot/apps/honeypot_server/global"
	"Honeypot/apps/honeypot_server/middleware"
	"Honeypot/apps/honeypot_server/models"
	"Honeypot/apps/honeypot_server/service/log_service"
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
	//log := middleware.GetLog(c)
	loginLog := log_service.NewLoginLog(c)
	if cr.CaptchaID == "" || cr.CaptchaCode == "" {
		loginLog.FailLog(cr.Username, "", "未输入图片验证码")
		resp.FailWithMsg("请输入图片验证码", c)
		return
	}
	if !captcha.CaptchaStore.Verify(cr.CaptchaID, cr.CaptchaCode, true) {
		loginLog.FailLog(cr.Username, "", "图片验证码验证失败")
		resp.FailWithMsg("图片验证码验证失败", c)
		return
	}

	var user models.UserModel
	err := global.DB.Take(&user, "username = ?", cr.Username).Error
	if err != nil {
		loginLog.FailLog(cr.Username, cr.Password, "用户名不存在")
		resp.FailWithMsg("用户名或密码错误", c)
		return
	}

	if !pwd.CompareHashAndPassword(user.Password, cr.Password) {
		loginLog.FailLog(cr.Username, cr.Password, "密码错误")
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

	loginLog.SuccessLog(user.ID, cr.Username)
	resp.OkWithData(token, c)
}
