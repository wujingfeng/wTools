package util

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	// 发送请求
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	// 读取内容
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
