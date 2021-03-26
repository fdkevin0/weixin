package base

import (
	"WeChat-golang/mp/core"
)

func GetCallbackIP(clt *core.Client) (ipList []string, err error) {
	var response struct {
		List []string `json:"ip_list"`
	}
	_, err = clt.GetJson("/cgi-bin/getcallbackip",
		nil,
		response,
	)
	if err != nil {
		ipList = response.List
	}
	return
}
