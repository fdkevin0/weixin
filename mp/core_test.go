package mp

import (
	"WeChat-golang/mp/core"
	"WeChat-golang/mp/user"
	"fmt"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	srv := &core.AccessTokenServer{
		AppId:     "wx5d769a50fe0a6589",
		AppSecret: "e87a13b7b2e4b39ac884eae60b004404",
	}
	srv.GetAccessToken()
	data, err := user.GetUserList(core.GetClient())
	fmt.Println(data.Openid[0])
	err = user.GetBaseInfo(core.GetClient(), data.Openid[0])
	if err != nil {
		fmt.Println(err)
		return
	}
}
