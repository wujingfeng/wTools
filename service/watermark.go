package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"wTools/enum/ten"
	"wTools/global"
	"wTools/param"
	"wTools/response"
	"wTools/util"
)

type Watermark struct {
	Code  int64  `json:"code"`
	Title string `json:"title"`
	Cover string `json:"cover"`
	Url   string `json:"url"`
}

func GetResponseParams(url string) (Watermark, error) {
	result := Watermark{}
	res, err := util.Get(url)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(res), &result)
	return result, err
}

func VideoDownload(url string, route string, ctx *gin.Context) {
	md5Str := util.Md5(url)
	fp := util.RootPath() + "/runtime/video/" + md5Str + "/"

	key := ten.VIDEO_KEY + md5Str
	//global.Redis.Del(key)
	fpr := global.Redis.Get(key)
	var err error
	filename := ""
	filepath := fpr.Val()
	if filepath != "" {
		fList := strings.Split(filepath, "/")
		filename = fList[len(fList)-1]
		if _, err = os.Stat(filepath); err == nil {
			// 如果文件存在, 则直接下载
			util.DownloadToClient(filepath, filename, ctx)
			return
		}

	}

	// 其他情况走直接下载流程
	url = ten.TEN_HOST + route + "?url=" + url
	res, _ := GetResponseParams(url)
	filename = res.Title + ".mp4"
	//res := Watermark{
	//	Code: 200,
	//	Title: "标题交给你们了。",
	//	Cover: "https://p3-ppx.byteimg.com/img/mosaic-legacy/2ab8400068ec7576befea~272x480.webp",
	//	Url: "http://v3-dy-o.zjcdn.com/ff5570374c30ce9d4f2ca0caf5fd2695/62a9abf1/video/tos/cn/tos-cn-ve-0076/41064fc495f04f029e8629421b1352fd/?a=1319&ch=0&cr=0&dr=3&cd=0%7C0%7C0%7C0&br=347&bt=347&cs=0&ds=1&ft=ArYleBnZqI2mo0PMpgafkVQR3MO-_KJ&mime_type=video_mp4&qs=0&rc=OTczZTxkNjo6NjplOTRpaEBpanl4PGd0bDl4bjMzZGYzM0BeYy8vLmMwNTMxYzIuNmFeYSMtYGsucWdjNDVfLS0yMS9zcw%3D%3D&l=202206151652360102091701630D4A2FFB",
	//}
	err = util.Download(fp, res.Url, filename, ctx)

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	global.Redis.Set(key, fp+filename, -1)

}

// PPX
//  @Description: 皮皮虾视频去水印
//  @param param
//  @param ctx
//  @author	wujingfeng 2022-06-15 17:06:59
func PPX(param param.ParamUrl, ctx *gin.Context) {
	VideoDownload(param.Url, ten.PPX, ctx)
}

// DouYin
//  @Description: 抖音
//  @param param
//  @param ctx
//  @author	wujingfeng 2022-06-15 22:28:22
func DouYin(param param.ParamUrl, ctx *gin.Context) {
	VideoDownload(param.Url, ten.DOUYIN, ctx)
}

// Video
//  @Description: 聚合接口
// 支持列表：抖音/皮皮虾/火山/微视/微博/绿洲/最右/轻视频/instagram/哔哩哔哩/快手/全民小视频/皮皮搞笑
// 全民k歌/巴塞电影/陌陌/Before避风/开眼/Vue Vlog/小咖秀/西瓜视频/逗拍
//  @param param
//  @param ctx
//  @author	wujingfeng 2022-06-15 22:41:28
func Video(param param.ParamUrl, ctx *gin.Context) {
	VideoDownload(param.Url, ten.VIDEO, ctx)
}

func ZuiYou(param param.ParamUrl, ctx *gin.Context) {
	VideoDownload(param.Url, ten.ZUIYOU, ctx)
}

func WeiShi(param param.ParamUrl, ctx *gin.Context) {
	VideoDownload(param.Url, ten.WEISHI, ctx)
}

func HuoShan(param param.ParamUrl, ctx *gin.Context) {
	VideoDownload(param.Url, ten.HUOSHAN, ctx)
}

func WeiBo(param param.ParamUrl, ctx *gin.Context) {
	VideoDownload(param.Url, ten.WEIBO, ctx)
}

func KuaiShou(param param.ParamUrl, ctx *gin.Context) {
	VideoDownload(param.Url, ten.KUAISHOU, ctx)
}

func QuanMin(param param.ParamUrl, ctx *gin.Context) {
	VideoDownload(param.Url, ten.QUANMIN, ctx)
}

func Ppgx(param param.ParamUrl, ctx *gin.Context) {
	VideoDownload(param.Url, ten.PPGX, ctx)
}

func XiGua(param param.ParamUrl, ctx *gin.Context) {
	VideoDownload(param.Url, ten.XIGUA, ctx)
}

func MoMo(param param.ParamUrl, ctx *gin.Context) {
	VideoDownload(param.Url, ten.MOMO, ctx)
}

func DouPai(param param.ParamUrl, ctx *gin.Context) {
	VideoDownload(param.Url, ten.DOUPAI, ctx)
}
