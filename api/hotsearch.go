package api

import (
	"github.com/gin-gonic/gin"
	"wTools/response"
	"wTools/service"
)

type HotSearch struct{}

func (hs *HotSearch) WeiBo(ctx *gin.Context) {
	HotSearchService := service.HotSearch{}
	response.OKWithData(HotSearchService.WeiBo(ctx), ctx)
	return
}

func (hs *HotSearch) DouYin(ctx *gin.Context) {
	HotSearchService := service.HotSearch{}
	response.OKWithData(HotSearchService.DouYin(ctx), ctx)
	return
}

func (hs *HotSearch) ZhiHu(ctx *gin.Context) {
	HotSearchService := service.HotSearch{}
	response.OKWithData(HotSearchService.ZhiHu(ctx), ctx)
	return
}
