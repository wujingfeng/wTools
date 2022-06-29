package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wTools/request"
	"wTools/response"
	"wTools/service"
)

type Watermark struct{}

func (wm *Watermark) Download(ctx *gin.Context) {
	//p := ctx.PostForm("path")
	p := ctx.Query("path")
	//p = "runtime/video/cf46d017a7a256861eb6d822e7a96ba0/标题交给你们了。.mp4"
	//p = "runtime/video/4f66561e249179123044dbac8591eac6/泪目！长安街上女学生下车敬礼，这一个动作，胜过千言！.mp4"
	fmt.Println(p)
	ctx.File(p)
	return
}

// Video
//  @Description: 聚合去水印, 自动识别地址来源
//  @param ctx
//  @author	wujingfeng 2022-06-22 10:31:08
func (wm *Watermark) Video(ctx *gin.Context) {
	param := request.ParamUrl(ctx)

	watermark := service.Watermark{}
	wms := watermark.Video(param, ctx)
	response.OKWithData(wms, ctx)
}

func (wm *Watermark) PPX(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	wms := watermark.PPX(param, ctx)
	response.OKWithData(wms, ctx)
}

func (wm *Watermark) DouYin(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	wms := watermark.DouYin(param, ctx)
	response.OKWithData(wms, ctx)
}

func (wm *Watermark) ZuiYou(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	wms := watermark.ZuiYou(param, ctx)
	response.OKWithData(wms, ctx)
}

func (wm *Watermark) WeiShi(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	wms := watermark.WeiShi(param, ctx)
	response.OKWithData(wms, ctx)
}

func (wm *Watermark) HuoShan(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	wms := watermark.HuoShan(param, ctx)
	response.OKWithData(wms, ctx)
}

func (wm *Watermark) WeiBo(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	wms := watermark.WeiBo(param, ctx)
	response.OKWithData(wms, ctx)
}

func (wm *Watermark) KuaiShou(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	wms := watermark.KuaiShou(param, ctx)
	response.OKWithData(wms, ctx)
}

func (wm *Watermark) QuanMin(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	wms := watermark.QuanMin(param, ctx)
	response.OKWithData(wms, ctx)
}

func (wm *Watermark) Ppgx(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	wms := watermark.Ppgx(param, ctx)
	response.OKWithData(wms, ctx)
}

func (wm *Watermark) XiGua(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	wms := watermark.XiGua(param, ctx)
	response.OKWithData(wms, ctx)
}
func (wm *Watermark) MoMo(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	wms := watermark.MoMo(param, ctx)
	response.OKWithData(wms, ctx)
}
func (wm *Watermark) DouPai(ctx *gin.Context) {
	param := request.ParamUrl(ctx)
	watermark := service.Watermark{}
	wms := watermark.DouPai(param, ctx)
	response.OKWithData(wms, ctx)
}
