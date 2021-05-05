package core

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/url"
)

type AccessTokenServer struct {
	AppKey    string `json:"appid"`  //腾讯微卡分配的授权凭证
	AppSecret string `json:"secret"` //授权凭证密钥
	OCode     string `json:"ocode"`  //主体代码
}

type accessTokenPostBody struct {
	AppKey    string `json:"app_key"`    //腾讯微卡分配的授权凭证
	AppSecret string `json:"app_secret"` //授权凭证密钥
	GrandType string `json:"grand_type"` //默认填：client_credentials
	Scope     string `json:"scope"`      //默认填：base
	OCode     string `json:"ocode"`      //主体代码
}

func (srv *AccessTokenServer) GetAccessToken() {
	//https://wiki.weixiao.qq.com/api/school/accessToken.html
	requestUrl := url.URL{
		Scheme: "https",
		Host:   "open.wecard.qq.com",
		Path:   "/cgi-bin/oauth2/token",
	}
	postBody := accessTokenPostBody{
		AppKey:    srv.AppKey,
		AppSecret: srv.AppSecret,
		GrandType: "client_credentials",
		Scope:     "base",
		OCode:     srv.OCode,
	}
	resp := new(AccessTokenResponse)
	gorequest.New().
		Post(requestUrl.String()).
		Send(postBody).
		EndStruct(resp)
	fmt.Println(resp.AccessToken, resp.ExpiresIn)
	client.AccessToken = resp.AccessToken
}
