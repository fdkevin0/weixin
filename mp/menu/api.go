package menu

import "WeChat-golang/mp/core"

func CreateMenu(clt core.Client) (err error) {
	_, err = clt.PostJson("/cgi-bin/menu/create",
		nil,
		nil,
		nil,
	)
	return
}

func GetMenu(clt core.Client) (err error) {
	_, err = clt.GetJson("/cgi-bin/get_current_selfmenu_info",
		nil,
		nil,
	)
	return
}
