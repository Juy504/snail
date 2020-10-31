package service

import (
	"encoding/json"
	"errors"
	d3auth "github.com/zcshan/d3outh"
	"regexp"
	common "snail/util"
	"strconv"
)
type AuthWx struct {

}

// 生成登录地址
func (w *AuthWx) GenUrl(state string) string{
	return ""
}

// 获取token信息
func (w *AuthWx) GetToken(code string) (*d3auth.AuthWxSucc, error){
	accessTokenUrl := "https://api.weixin.qq.com/sns/oauth2/access_token?appid="
	str, err := common.HttpGet(accessTokenUrl)
	if err != nil{
		return nil, err
	}
	ismatch, _ := regexp.MatchString("errcode", str)
	if ismatch {
		p := &d3auth.AuthWxErr{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("Error:" + strconv.Itoa(p.Error) + " Error_description:" + p.ErrorDescription)
	} else {
		p := &d3auth.AuthWxSucc{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
	}
	return &d3auth.AuthWxSucc{}, nil
}
