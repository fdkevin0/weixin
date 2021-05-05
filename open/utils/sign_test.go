package utils

import (
	"WeChat-golang/weixiao/core"
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {
	var correctSign = "5FC8D3BC9EE19B1DCAC067126D0B83E3"
	srv := core.Client{
		AppInfo: core.AccessTokenServer{
			AppSecret: "CF8CDD8D0B7B1ACDA76201F406BED81E",
		},
	}
	s := Sign(srv, map[string]string{
		"media_id":  "gh_2s97120599f5",
		"api_key":   "BE30B2BA30ABFADB",
		"timestamp": "1442401156",
		"nonce_str": "9A0A8659F005D6984697E2CA0A9CF399",
	})
	fmt.Println(s)
	if s != correctSign {
		panic("sign error")
	}
}
