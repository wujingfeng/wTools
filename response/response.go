package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = 1
	ERROR   = 2
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Result(code int, msg string, data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// OK
//  @Description: 成功
//  @receiver res
//  @param ctx
//  @author	wujingfeng 2022-06-15 10:43:13
func OK(ctx *gin.Context) {
	Result(SUCCESS, "操作成功", "", ctx)
}

func OKWithMessage(msg string, ctx *gin.Context) {
	Result(SUCCESS, msg, "", ctx)
}

func OKWithData(data interface{}, ctx *gin.Context) {
	Result(SUCCESS, "操作成功", data, ctx)
}

// Fail
//  @Description: 失败
//  @receiver res
//  @param ctx
//  @author	wujingfeng 2022-06-15 10:43:40
func Fail(ctx *gin.Context) {
	Result(ERROR, "操作失败", "", ctx)
}

func FailWithMessage(msg string, ctx *gin.Context) {
	Result(ERROR, msg, "", ctx)
}

func FailWithData(data interface{}, ctx *gin.Context) {
	Result(ERROR, "操作失败", data, ctx)
}
