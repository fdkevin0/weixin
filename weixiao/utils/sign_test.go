package utils

import (
	"WeChat-golang/weixiao/core"
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {
	srv := core.Client{
		AppInfo: core.AccessTokenServer{
			AppSecret: "12345678123456781234567812345678",
		},
	}
	s := Sign(srv, map[string]string{
		"app_key":     "123",
		"timestamp":   "123",
		"nonce":       "123",
		"school_code": "123",
	})
	fmt.Println(s)
}
