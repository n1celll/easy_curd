package wechat

import (
	"easy_curd/conf"
	"fmt"
	"github.com/imroc/req"
)

type TokenInfo struct {
	Access_token string `json:"access_token"`
	Openid       string `json:"openid"`
	Unionid      string `json:"unionid"`
}

type WechatUserInfo struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Unionid  string `json:"unionid"`
	Openid   string `json:"openid"`
}

func getAccessToken(code string) *TokenInfo {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		conf.WX_APPID, conf.WX_APPSECRET, code)
	r, err := req.Get(url)
	if err != nil {
		return nil
	}
	tokenInfo := &TokenInfo{}
	err = r.ToJSON(tokenInfo)
	if err != nil {
		return nil
	}
	return tokenInfo
}

func GetUserInfo(code string) (*WechatUserInfo, error) {
	tokenInfo := getAccessToken(code)
	if tokenInfo == nil {
		return nil, nil
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s",
		tokenInfo.Access_token, tokenInfo.Openid)
	userInfo := &WechatUserInfo{}

	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	err = r.ToJSON(userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
