package identify

import (
	"WeChat-golang/weixiao/core"
	"github.com/parnurzeal/gorequest"
	"net/url"
	"time"
)

func GetUserDataByCampusCode(clt core.Client, scene int, nonce, deviceNo, location, authCode string) {
	//https://wiki.weixiao.qq.com/api/school/campuscode.html
	requestUrl := url.URL{
		Scheme: "https",
		Host:   "weixiao.qq.com",
		Path:   "/apps/school-api/campus-code",
	}
	postBody := campusCodePostBody{
		AppKey:     clt.AppInfo.AppKey,
		Timestamp:  time.Now().Unix(),
		Nonce:      nonce,
		Scene:      scene,
		DeviceNo:   deviceNo,
		Location:   location,
		AuthCode:   authCode,
		SchoolCode: clt.AppInfo.OCode,
	}

	resp := new(campusCodeResponse)
	gorequest.New().
		Post(requestUrl.String()).
		Send(postBody).
		EndStruct(resp)
}
