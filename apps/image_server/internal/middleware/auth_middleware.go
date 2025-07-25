// Package middleware 权限验证
package middleware

import (
	"Honeypot/apps/image_server/internal/enum"
	"Honeypot/apps/image_server/internal/global"
	"Honeypot/apps/image_server/internal/utils/jwts"
	"Honeypot/apps/image_server/internal/utils/resp"
	"Honeypot/apps/image_server/internal/utils/white_list"
	"fmt"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware 权限验证
func AuthMiddleware(c *gin.Context) {
	// 去判断这个路径在不在白名单中
	path := c.Request.URL.Path
	fmt.Println("white_list:", global.Config.WhiteList)
	fmt.Println("path:", path)
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
