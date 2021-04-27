package menu

import (
	"WeChat-golang/mp/core"
	"WeChat-golang/mp/menu/createMenu"
	"WeChat-golang/mp/menu/getMenu"
	"fmt"
)

func Create(clt core.Client) (err error) {
	// https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Creating_Custom-Defined_Menu.html
	_, err = clt.PostJson("/cgi-bin/menu/create",
		nil,
		&createMenu.PostBody{
			Button: []createMenu.Button{
				{
					Type: "click",
					Name: "今日歌曲",
					Key:  "V1001_TODAY_MUSIC",
				},
			},
		},
		nil,
	)
	return
}

func Get(clt core.Client) (err error) {
	//https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Querying_Custom_Menus.html
	resp := new(getMenu.Response)
	_, err = clt.GetJson("/cgi-bin/get_current_selfmenu_info",
		nil,
		resp,
	)
	fmt.Println(resp.SelfMenuInfo.Button)
	return
}
