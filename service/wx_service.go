package service

import (
	"snail/conf"
	"encoding/json"
	"errors"
	"regexp"
	"snail/util"
	"strconv"
)

// 生成登录地址
func (w *AuthWx) GenUrl(state string) string{
	return ""
}

type AuthWx struct {
	Conf *conf.Wxconf
}

type AuthWxErr struct {
	Error             int    `json:"errcode"`
	ErrorDescription string `json:"errmsg"`
}

type AuthWxSucc struct {
	AccessToken string `json:"access_token"`
	Openid       string `json:"openid"`
}


// 获取token信息
func (w *AuthWx) GetToken(code string) (*AuthWxSucc, error){
	accessTokenUrl := "https://api.weixin.qq.com/sns/oauth2/access_token?appid="
	str, err := util.HttpGet(accessTokenUrl)
	if err != nil{
		return nil, err
	}
	ismatch, _ := regexp.MatchString("errcode", str)
	if ismatch {
		p := &AuthWxErr{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("Error:" + strconv.Itoa(p.Error) + " Error_description:" + p.ErrorDescription)
	} else {
		p := &AuthWxSucc{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
	}
	return &AuthWxSucc{}, nil
}
