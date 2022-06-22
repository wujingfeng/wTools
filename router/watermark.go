package router

import (
	"github.com/gin-gonic/gin"
	"wTools/api"
)

func InitWatermark(router *gin.RouterGroup) {
	watermark := api.Watermark{}
	r := router.Group("/watermark")
	{
		r.GET("video", watermark.Video)
		r.GET("ppx", watermark.PPX)
		r.GET("douyin", watermark.DouYin)
		r.GET("zuiyou", watermark.ZuiYou)
		r.GET("weishi", watermark.WeiShi)
		r.GET("huoshan", watermark.HuoShan)
		r.GET("weibo", watermark.WeiBo)
		r.GET("kuaishou", watermark.KuaiShou)
		r.GET("quanmin", watermark.QuanMin)
		r.GET("ppgx", watermark.Ppgx)
		r.GET("xigua", watermark.XiGua)
		r.GET("momo", watermark.MoMo)
		r.GET("doupai", watermark.DouPai)
	}
}
