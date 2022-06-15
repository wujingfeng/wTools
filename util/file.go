package util

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Download(filepath string, url string, filename string, ctx *gin.Context) error {
	err := DownloadToServer(filepath, filename, url)
	if err != nil {
		return err
	}
	DownloadToClient(filepath+filename, filename, ctx)
	return nil
}

// DownloadToServer
//  @Description: 下载视频到服务器
//  @param filepath	string 指定下载位置
//  @param url		string 远程地址
//  @return error
//  @author	wujingfeng 2022-06-15 11:06:54
func DownloadToServer(filepath string, filename string, url string) error {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer res.Body.Close()

	err = os.MkdirAll(filepath, 0755)
	if err != nil {
		return err
	}

	out, err := os.Create(filepath + filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, res.Body)
	return err
}

// DownloadToClient
//  @Description: 文件下载到客户端
//  @param filepath
//  @param filename
//  @param ctx
//  @author	wujingfeng 2022-06-15 16:59:30
func DownloadToClient(filepath string, filename string, ctx *gin.Context) {
	ctx.Header("Content-Length", "-1")
	ctx.Header("Transfer-Encoding", "true")
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	// 传输二进制流到浏览器
	ctx.File(filepath)
}

// RootPath
//  @Description: 获取项目根目录
//  @return string
//  @author	wujingfeng 2022-06-15 16:34:32
func RootPath() string {
	rootPath, err := filepath.Abs("")
	if err != nil {
		panic(err)
	}
	return rootPath

	//	et, err := os.Executable()
	//	if err != nil {
	//		panic(err)
	//	}
	//	wd, err := os.Getwd()
	//	fmt.Println(wd)
	//	fb,err := filepath.Abs("")
	//fmt.Println(fb)
	//	etPath := filepath.Dir(et)
	//	realPath, err := filepath.EvalSymlinks(etPath)
	//	if err != nil {
	//		panic(err)
	//	}
	//	return filepath.Dir(realPath)
}

// Md5
//  @Description: 字符串 md5加密
//  @param string
//  @return string
//  @author	wujingfeng 2022-06-15 17:25:53
func Md5(string string) string {
	hash := md5.Sum([]byte(string))
	md5Str := fmt.Sprintf("%x", hash)
	return md5Str
}
