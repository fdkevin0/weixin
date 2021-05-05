package base

type getInfoPostBody struct {
	MediaId   string //公众号原始ID
	ApiKey    string //应用KEY
	Timestamp int64  //时间戳
	NonceStr  string //32位随机字符串
	Sign      string //按照签名算法生成的签名（参考签名算法）
}

type getInfoResponse struct {
	Name        string `json:"name"`         //公众号名称
	MediaNumber string `json:"media_number"` //公众号设置的微信号（可能为空）
	AvatarImage string `json:"avatar_image"` //公众号头像
	MediaID     string `json:"media_id"`     //公众号原始ID
	//可通过微信的接口带上公众号原始ID参数获取二维码
	//Example：https://open.weixin.qq.com/qr/code?username=gh_560e659c5877
	MediaType  string `json:"media_type"`  //公众号类型（ 0代表订阅号，1代表由历史老帐号升级后的订阅号，2代表服务号）
	MediaURL   string `json:"media_url"`   //公众号信息页（展示公众号基本信息和近期推文，可进行快捷关注。由于部分公众号信息不完整，该字段可能为空）
	SchoolName string `json:"school_name"` //公众号学校/企业名称（非高校公众号或者学校/企业审核不通过的为空）
	SchoolCode string `json:"school_code"` //学校/企业标识码（非高校公众号或者学校/企业审核不通过的为空）
	//查看本校/企业代码 http://www.moe.gov.cn/srcsite/A03/moe_634/201706/t20170614_306900.html
	VerifyType string `json:"verify_type"` //公众号认证类型
	//-1 = "未认证"
	//0 = "微信认证"
	//1 = "新浪微博认证"
	//2 = "腾讯微博认证"
	//3 = "已资质认证通过但还未通过名称认证"
	//4 = "已资质认证通过、还未通过名称认证，但通过了新浪微博认证"
	//5 = "已资质认证通过、还未通过名称认证，但通过了腾讯微博认证"
}
