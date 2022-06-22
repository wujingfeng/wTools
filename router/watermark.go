package router

import (
	"github.com/gin-gonic/gin"
	"wTools/api"
)

func InitWatermark(router *gin.RouterGroup) {
	r := router.Group("/watermark")
	{
		r.GET("video", (api.Watermark{}).Video)
		r.GET("ppx", (api.Watermark{}).PPX)
		r.GET("douyin", (api.Watermark{}).DouYin)
		r.GET("zuiyou", (api.Watermark{}).ZuiYou)
		r.GET("weishi", (api.Watermark{}).WeiShi)
		r.GET("huoshan", (api.Watermark{}).HuoShan)
		r.GET("weibo", (api.Watermark{}).WeiBo)
		r.GET("kuaishou", (api.Watermark{}).KuaiShou)
		r.GET("quanmin", (api.Watermark{}).QuanMin)
		r.GET("ppgx", (api.Watermark{}).Ppgx)
		r.GET("xigua", (api.Watermark{}).XiGua)
		r.GET("momo", (api.Watermark{}).MoMo)
		r.GET("doupai", (api.Watermark{}).DouPai)
	}
}
