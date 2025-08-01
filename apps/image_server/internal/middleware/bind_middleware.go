package middleware

import (
	"github.com/gin-gonic/gin"
	"image_server/internal/utils/resp"
)

func BindJsonMiddleware[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithError(err, c)
		c.Abort()
		return
	}
	c.Set("request", cr)
}

func BindQueryMiddleware[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		resp.FailWithError(err, c)
		c.Abort()
		return
	}
	c.Set("request", cr)
}

func BindUriMiddleware[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindUri(&cr)
	if err != nil {
		resp.FailWithError(err, c)
		c.Abort()
		return
	}
	c.Set("request", cr)
}

func GetBind[T any](c *gin.Context) (cr T) {
	return c.MustGet("request").(T)
}
