package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"wTools/enum/ten"
	"wTools/global"
	"wTools/response"
	"wTools/util"
)

type HotSearch struct {
	Data int64       `json:"data"`
	List interface{} `json:"list"`
}

// 微博

type HotSearchWeibo struct {
	HotSearch `json:"-"`
	Data      int64   `json:"data"`
	List      []weibo `json:"list"`
}
type weibo struct {
	Name string `json:"name"`
	Hot  int64  `json:"hot"`
	Url  string `json:"url"`
}

// 知乎

type HotSearchZhiHu struct {
	HotSearch `json:"-"`
	Data      int64   `json:"data"`
	List      []zhihu `json:"list"`
}
type zhihu struct {
	Name  string `json:"name"`
	Query string `json:"query"`
	Url   string `json:"url"`
}

// 抖音

type HotSearchDouYin struct {
	HotSearch `json:"-"`
	Data      int64    `json:"data"`
	List      []douyin `json:"list"`
}
type douyin struct {
	Name string `json:"name"`
	Hot  string `json:"hot"`
}

func (hs *HotSearch) search(route string, model interface{}) (interface{}, error) {
	url := ten.TEN_HOST + route

	res, err := util.Get(url)
	if err != nil {
		fmt.Println("err:" + err.Error())
		return model, err
	}
	err = json.Unmarshal([]byte(res), &model)
	return model, err

}

// WeiBo
//  @Description: 微博热搜
//  @receiver hs
//  @param ctx
//  @return []weibo
//  @author	wujingfeng 2022-06-22 11:21:14
func (hs *HotSearch) WeiBo(ctx *gin.Context) []weibo {
	epr1 := util.TodayLastSeconds()
	fmt.Println(epr1)

	today := util.Today()
	res := global.Redis.Get(ten.HOT_SEARCH_KEY_WEIBO + today)
	wb := []weibo{}
	if res.Val() != "" {
		err := json.Unmarshal([]byte(res.Val()), &wb)
		if err == nil {
			// 缓存有效则直接返回
			return wb
		}
	}

	hotSearchWeibo := HotSearchWeibo{}
	result, err := hs.search(ten.HOT_SEARCH_WEIBO, hotSearchWeibo)
	if err != nil {
		response.Fail(ctx)
		return hotSearchWeibo.List
	}
	//hotSearchWeibo,ok := result.(HotSearchWeibo)
	resultByte, err := json.Marshal(result)
	err = json.Unmarshal(resultByte, &hotSearchWeibo)
	if err != nil {
		return []weibo{}
	}
	// 设置缓存
	wbStr, err := json.Marshal(hotSearchWeibo.List)
	epr := util.TodayLastSeconds()
	global.Redis.Set(ten.HOT_SEARCH_KEY_WEIBO+today, wbStr, time.Second*time.Duration(epr))
	return hotSearchWeibo.List
}

// DouYin
//  @Description: 抖音热搜
//  @receiver hs
//  @param ctx
//  @return []douyin
//  @author	wujingfeng 2022-06-22 14:51:07
func (hs *HotSearch) DouYin(ctx *gin.Context) []douyin {
	today := util.Today()
	res := global.Redis.Get(ten.HOT_SEARCH_KEY_DouYin + today)
	wb := []douyin{}
	if res.Val() != "" {
		err := json.Unmarshal([]byte(res.Val()), &wb)
		if err == nil {
			// 缓存有效则直接返回
			return wb
		}
	}

	hotSearchDouYin := HotSearchDouYin{}
	result, err := hs.search(ten.HOT_SEARCH_DOUYIN, hotSearchDouYin)
	if err != nil {
		response.Fail(ctx)
		return hotSearchDouYin.List
	}
	//hotSearchDouYin,ok := result.(hotSearchDouYin)
	resultByte, err := json.Marshal(result)
	err = json.Unmarshal(resultByte, &hotSearchDouYin)
	if err != nil {
		return []douyin{}
	}
	// 设置缓存
	wbStr, err := json.Marshal(hotSearchDouYin.List)
	epr := util.TodayLastSeconds()
	global.Redis.Set(ten.HOT_SEARCH_KEY_DouYin+today, wbStr, time.Second*time.Duration(epr))
	return hotSearchDouYin.List
}

// ZhiHu
//  @Description: 知乎热搜
//  @receiver hs
//  @param ctx
//  @return []zhihu
//  @author	wujingfeng 2022-06-22 14:50:57
func (hs *HotSearch) ZhiHu(ctx *gin.Context) []zhihu {
	today := util.Today()
	res := global.Redis.Get(ten.HOT_SEARCH_KEY_ZhiHu + today)
	wb := []zhihu{}
	if res.Val() != "" {
		err := json.Unmarshal([]byte(res.Val()), &wb)
		if err == nil {
			// 缓存有效则直接返回
			return wb
		}
	}

	hotSearchZhiHu := HotSearchZhiHu{}
	result, err := hs.search(ten.HOT_SEARCH_ZHIHU, hotSearchZhiHu)
	if err != nil {
		response.Fail(ctx)
		return hotSearchZhiHu.List
	}
	//hotSearchZhiHu,ok := result.(hotSearchZhiHu)
	resultByte, err := json.Marshal(result)
	err = json.Unmarshal(resultByte, &hotSearchZhiHu)
	if err != nil {
		return []zhihu{}
	}
	// 设置缓存
	wbStr, err := json.Marshal(hotSearchZhiHu.List)
	epr := util.TodayLastSeconds()
	global.Redis.Set(ten.HOT_SEARCH_KEY_DouYin+today, wbStr, time.Second*time.Duration(epr))
	return hotSearchZhiHu.List
}
