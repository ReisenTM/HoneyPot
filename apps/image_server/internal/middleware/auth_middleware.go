// Package middleware 权限验证
package middleware

import (
	"github.com/gin-gonic/gin"
	"image_server/internal/enum"
	"image_server/internal/global"
	"image_server/internal/utils/jwts"
	"image_server/internal/utils/resp"
	"image_server/internal/utils/white_list"
)

// AuthMiddleware 权限验证
func AuthMiddleware(c *gin.Context) {
	// 去判断这个路径在不在白名单中
	path := c.Request.URL.Path

	if white_list.WhiteListCheck(global.Config.WhiteList, path) {
		// 在白名单中，直接放行
		c.Next()
		return
	}
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		resp.FailWithError(err, c)
		c.Abort()
		return
	}

	//保存验证过的用户信息
	c.Set("claims", claims)
	return
}

// AdminMiddleware 管理员级验证
func AdminMiddleware(c *gin.Context) {

	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		resp.FailWithError(err, c)
		c.Abort()
		return
	}

	if claims.Role != enum.RoleAdmin {
		//不是管理员
		resp.FailWithMsg("权限错误", c)
		c.Abort()
		return
	}
	//保存验证过的用户信息
	c.Set("claims", claims)
	return
}

func GetAuth(c *gin.Context) *jwts.MyClaims {
	return c.MustGet("claims").(*jwts.MyClaims)
}
