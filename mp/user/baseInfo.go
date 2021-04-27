package user

import (
	"WeChat-golang/mp/core"
	"net/url"
)

type BaseInfo struct {
	Subscribe      int64   `json:"subscribe"`
	Openid         string  `json:"openid"`
	Nickname       string  `json:"nickname"`
	Sex            int64   `json:"sex"`
	Language       string  `json:"language"`
	City           string  `json:"city"`
	Province       string  `json:"province"`
	Country        string  `json:"country"`
	HeadImgUrl     string  `json:"headimgurl"`
	SubscribeTime  int64   `json:"subscribe_time"`
	UnionId        string  `json:"unionid"`
	Remark         string  `json:"remark"`
	GroupId        int64   `json:"groupid"`
	TagIdList      []int64 `json:"tagid_list"`
	SubscribeScene string  `json:"subscribe_scene"`
	QrScene        int64   `json:"qr_scene"`
	QrSceneStr     string  `json:"qr_scene_str"`
}

func GetBaseInfo(clt core.Client, openId string) (err error) {
	//https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
	resp := new(BaseInfo)
	_, err = clt.GetJson("/cgi-bin/user/info",
		url.Values{
			"openid": {openId},
			"lang":   {"zh_CN"},
		},
		resp,
	)
	return
}
