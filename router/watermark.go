package router

import (
	"github.com/gin-gonic/gin"
	"wTools/api"
)

func InitWatermark(router *gin.RouterGroup) {
	r := router.Group("/watermark")
	{
		r.GET("video", api.Video)
		r.GET("ppx", api.PPX)
		r.GET("douyin", api.DouYin)
		r.GET("zuiyou", api.ZuiYou)
		r.GET("weishi", api.WeiShi)
		r.GET("huoshan", api.HuoShan)
		r.GET("weibo", api.WeiBo)
		r.GET("kuaishou", api.KuaiShou)
		r.GET("quanmin", api.QuanMin)
		r.GET("ppgx", api.Ppgx)
		r.GET("xigua", api.XiGua)
		r.GET("momo", api.MoMo)
		r.GET("doupai", api.DouPai)
	}
}
