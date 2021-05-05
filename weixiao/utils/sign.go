package utils

import (
	"WeChat-golang/weixiao/core"
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

func Sign(srv core.Client, param map[string]string) string {
	//https://wiki.weixiao.qq.com/api/school/signAlgorithm.html
	//签名验证地址 https://weixiao.qq.com/public/signUtil/index.html
	//参数名ASCII码从小到大排序（字典序）
	var paramKey []string
	for k := range param {
		paramKey = append(paramKey, k)
	}
	sort.Strings(paramKey)
	paramArray := url.Values{}
	for _, v := range paramKey {
		if param[v] != "" { //如果参数的值为空不参与签名
			paramArray.Add(v, param[v])
		}
	}
	//在stringA最后拼接上&key=APP_SECRET得到stringSignTemp字符串
	paramArrayStr := paramArray.Encode() + "&key=" + srv.AppInfo.AppSecret
	//并对stringSignTemp进行MD5运算
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(paramArrayStr))
	//再将得到的字符串所有字符转换为大写，得到sign值signValue
	return strings.ToUpper(hex.EncodeToString(md5Ctx.Sum(nil)))
}
