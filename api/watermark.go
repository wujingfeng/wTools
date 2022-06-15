package api

import (
	"github.com/gin-gonic/gin"
	"wTools/request"
	"wTools/service"
)

func Video(ctx *gin.Context) {
	param := request.ParamUrl(ctx)

	service.Video(param, ctx)
}

func PPX(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	service.PPX(param, ctx)
}

func DouYin(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	service.DouYin(param, ctx)
}

func ZuiYou(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	service.ZuiYou(param, ctx)
}

func WeiShi(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	service.WeiShi(param, ctx)
}

func HuoShan(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	service.HuoShan(param, ctx)
}

func WeiBo(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	service.WeiBo(param, ctx)
}

func KuaiShou(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	service.KuaiShou(param, ctx)
}

func QuanMin(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	service.QuanMin(param, ctx)
}

func Ppgx(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	service.Ppgx(param, ctx)
}

func XiGua(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	service.XiGua(param, ctx)
}
func MoMo(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	service.MoMo(param, ctx)
}
func DouPai(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	service.DouPai(param, ctx)
}
