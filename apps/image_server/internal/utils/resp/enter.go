package resp

import (
	"Honeypot/apps/image_server/internal/utils/validate"
	"github.com/gin-gonic/gin"
	"net/http"
)

//响应封装

var empty = map[string]any{}

type Code int

func (c Code) String() string {
	switch c {
	case SuccessCode:
		return "成功"
	case FailCode:
		return "失败"
	default:
	}
	return ""
}

const (
	SuccessCode Code = iota
	FailCode
)

type Response struct {
	Code Code   `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func (r *Response) Json(c *gin.Context) {
	c.JSON(http.StatusOK, r)
}

func Ok(data any, msg string, c *gin.Context) {
	response := Response{Code: SuccessCode, Msg: msg, Data: data}
	response.Json(c)
}
func OkWithMsg(msg string, c *gin.Context) {
	response := Response{Code: SuccessCode, Msg: msg, Data: empty}
	response.Json(c)
}
func OkWithData(data any, c *gin.Context) {
	resp := Response{SuccessCode, data, "成功"}
	resp.Json(c)
}
func OkWithList(list any, count int64, c *gin.Context) {
	resp := Response{SuccessCode, map[string]any{
		"list":  list,
		"count": count,
	}, "成功"}
	resp.Json(c)
}
func FailWithCode(code Code, c *gin.Context) {
	response := Response{Code: code, Msg: "失败", Data: empty}
	response.Json(c)
}
func FailWithMsg(msg string, c *gin.Context) {
	response := Response{Code: FailCode, Msg: msg, Data: empty}
	response.Json(c)
}
func FailWithData(data any, c *gin.Context) {
	response := Response{
		Code: FailCode,
		Msg:  "失败",
		Data: data,
	}
	response.Json(c)
}
func FailWithError(err error, c *gin.Context) {
	msg := validate.ValidateError(err)
	response := Response{Code: FailCode, Msg: msg, Data: nil}
	response.Json(c)
}
