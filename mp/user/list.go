package user

import (
	"WeChat-golang/mp/core"
	"fmt"
	"net/url"
)

type GetUserListResponse struct {
	Total      int64  `json:"total"`
	Count      int64  `json:"count"`
	Data       Data   `json:"data"`
	NextOpenid string `json:"next_openid"`
}

type Data struct {
	Openid []string `json:"openid"`
}

func GetUserList(clt core.Client) (data Data, err error) {
	//https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
	resp := new(GetUserListResponse)
	_, err = clt.GetJson("/cgi-bin/user/get",
		url.Values{
			//"openid": {openId},
		},
		resp,
	)
	data = resp.Data
	fmt.Println(resp)
	return
}
