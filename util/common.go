package util

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

func HttpGet(url string)(string, error){
	trans := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// 创建http客户端
	// 创建cookie
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Transport: trans,
		Jar: cookieJar,
	}
	// 同步发送http请求
	resp, err := client.Get(url)
	if err != nil{
		return "", err
	}
	// 这里对body进行关闭
	// 1、源码中对此做了特别强调 2、Body本身其实是嵌套了多层的TCPConn为了连接复用，需要关闭
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return "", err
	}
	return string(body), nil
}

func HttpPost(url string)(string, error){
	trans := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: trans,
	}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil{
		return "", err
	}
	resp, _ := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return "", err
	}
	return string(body), nil
}
