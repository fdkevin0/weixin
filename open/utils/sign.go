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
	//https://open.weixiao.qq.com/app/index.html#/api?content=auth&_k=b0ph2q
	//签名验证地址 https://pay.weixin.qq.com/wiki/tools/signverify/
	//拼接集合键值对字符串
	//设所有数据为集合M，将集合M按照参数名ASCII码从小到大排序（字典序）
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
	//使用URL键值对的格式（即key1=value1&key2=value2…）拼接成字符串stringA
	//在stringA最后拼接上api_secret得到stringSignTemp字符串
	paramArrayStr := paramArray.Encode() + "&key=" + srv.AppInfo.AppSecret
	//并对stringSignTemp进行MD5运算
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(paramArrayStr))
	//再将得到的字符串所有字符转换为大写，得到sign值signValue
	return strings.ToUpper(hex.EncodeToString(md5Ctx.Sum(nil)))
}
