package subMessage

import "WeChat-golang/mp/core"

type addTemplatePostBody struct {
	Tid       string `json:"tid"`       //模板标题 id，可通过getPubTemplateTitleList接口获取，也可登录公众号后台查看获取
	KidList   []int  `json:"kidList"`   //开发者自行组合好的模板关键词列表，关键词顺序可以自由搭配 （例如 [3,5,4] 或 [4,5,3]），最多支持5个，最少2个关键词组合
	SceneDesc string `json:"sceneDesc"` //服务场景描述，15个字以内
}

type addTemplateResponse struct {
	core.Error
	PriTmplId string `json:"priTmplId"` //添加至帐号下的模板id，发送订阅通知时所需
}

type deleteTemplatePostBody struct {
	PriTmplId string `json:"priTmplId"` //要删除的模板id
}

type getCategoryResponse struct {
	*core.Error
	Data []struct {
		Id   int    `json:"id"`   //类目id，查询公共模板库时需要
		Name string `json:"name"` //类目的中文名
	}
}

type getPubTemplateKeyWordsByIdResponse struct {
	*core.Error
	Data []struct {
		Kid     int    `json:"kid"`     //关键词 id，选用模板时需要
		Name    string `json:"name"`    //关键词内容
		Example string `json:"example"` //关键词内容对应的示例
		Rule    string `json:"rule"`    //参数类型
	} `json:"data"` //关键词列表
}

type getPubTemplateTitleListResponse struct {
	*core.Error
	Count int `json:"count"` //公共模板列表总数
	Data  struct {
		Tid        int    `json:"tid"`        //模版标题 id
		Title      string `json:"title"`      //模版标题
		Type       int    `json:"type"`       //模版类型，2 为一次性订阅，3 为长期订阅
		CategoryId int    `json:"categoryId"` //模版所属类目 id
	} `json:"data"` //模板标题列表
}

type getTemplateListResponse struct {
	*core.Error
	Count int `json:"count"` //公共模板列表总数
	Data  struct {
		PriTmplId string `json:"priTmplId"` //添加至帐号下的模板 id，发送订阅通知时所需
		Title     string `json:"title"`     //模版标题
		Content   string `json:"content"`   //模版内容
		Example   string `json:"example"`   //模板内容示例
		Type      int    `json:"type"`      //模版类型，2 为一次性订阅，3 为长期订阅
	} `json:"data"` //个人模板列表
}

type sendPostBody struct {
	ToUser      string            `json:"touser"`      //接收者（用户）的 openid
	TemplateId  string            `json:"template_id"` //所需下发的订阅模板id
	Page        string            `json:"page"`        //跳转网页时填写
	MiniProgram *core.MiniProgram `json:"miniprogram"` //跳转小程序时填写，格式如{ "appid": "pagepath": { "value": any } }
	Data        interface{}       `json:"data"`        //模板内容，格式形如 { "key1": { "value": any }, "key2": { "value": any } }
}
