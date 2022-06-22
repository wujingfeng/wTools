package router

import (
	"github.com/gin-gonic/gin"
	"wTools/api"
)

func InitHotSearch(router *gin.RouterGroup) {
	hotSearch := api.HotSearch{}

	r := router.Group("hot")
	{
		r.GET("weibo", hotSearch.WeiBo)
		r.GET("zhihu", hotSearch.ZhiHu)
		r.GET("douyin", hotSearch.DouYin)
	}
}
