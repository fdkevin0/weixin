package core

import (
	"github.com/parnurzeal/gorequest"
	"net/url"
)

type AccessTokenServer struct {
	AppId     string `json:"appid"`  //第三方用户唯一凭证
	AppSecret string `json:"secret"` //第三方用户唯一凭证密钥，即appsecret
}

//https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
func (srv *AccessTokenServer) GetAccessToken() {
	requestUrl := url.URL{
		Scheme: "https",
		Host:   "api.weixin.qq.com",
		Path:   "/cgi-bin/token",
		RawQuery: url.Values{
			"grant_type": []string{"client_credential"}, //获取access_token填写client_credential
			"appid":      []string{srv.AppId},
			"secret":     []string{srv.AppSecret},
		}.Encode(),
	}
	gorequest.New().
		Get(requestUrl.String()).
		End()
}
