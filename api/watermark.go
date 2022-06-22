package api

import (
	"github.com/gin-gonic/gin"
	"wTools/request"
	"wTools/service"
)

type Watermark struct{}

// Video
//  @Description: 聚合去水印, 自动识别地址来源
//  @param ctx
//  @author	wujingfeng 2022-06-22 10:31:08
func (wm *Watermark) Video(ctx *gin.Context) {
	param := request.ParamUrl(ctx)

	watermark := service.Watermark{}
	watermark.Video(param, ctx)
}

func (wm *Watermark) PPX(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	watermark.PPX(param, ctx)
}

func (wm *Watermark) DouYin(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	watermark.DouYin(param, ctx)
}

func (wm *Watermark) ZuiYou(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	watermark.ZuiYou(param, ctx)
}

func (wm *Watermark) WeiShi(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	watermark.WeiShi(param, ctx)
}

func (wm *Watermark) HuoShan(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	watermark.HuoShan(param, ctx)
}

func (wm *Watermark) WeiBo(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	watermark.WeiBo(param, ctx)
}

func (wm *Watermark) KuaiShou(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	watermark.KuaiShou(param, ctx)
}

func (wm *Watermark) QuanMin(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	watermark.QuanMin(param, ctx)
}

func (wm *Watermark) Ppgx(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	watermark.Ppgx(param, ctx)
}

func (wm *Watermark) XiGua(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	watermark.XiGua(param, ctx)
}
func (wm *Watermark) MoMo(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	watermark.MoMo(param, ctx)
}
func (wm *Watermark) DouPai(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	watermark.DouPai(param, ctx)
}
