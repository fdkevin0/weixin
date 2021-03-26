package core

type Error struct {
	ErrorCode    int    `json:"errcode"` //错误码
	ErrorMessage string `json:"errmsg"`  //错误信息
}

type GetAccessTokenResponse struct {
	AccessToken string `json:"access_token"` //获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   //凭证有效时间，单位：秒
	*Error
}
