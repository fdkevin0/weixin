package subMessage

import (
	"WeChat-golang/mp/core"
	"net/url"
)

//选用模板
func AddTemplate(clt *core.Client, tid string, kidList []int, sceneDesc string) (priTmplId string, err error) {
	response := new(addTemplateResponse)
	_, err = clt.PostJson("/wxaapi/newtmpl/addtemplate",
		nil,
		addTemplatePostBody{
			Tid:       tid,
			KidList:   kidList,
			SceneDesc: sceneDesc,
		},
		response,
	)
	if err != nil {
		priTmplId = response.PriTmplId
	}
	return
}

//删除模板
func DeleteTemplate(clt *core.Client, priTmplId string) (err error) {
	_, err = clt.PostJson(
		"/wxaapi/newtmpl/deltemplate",
		nil,
		deleteTemplatePostBody{
			PriTmplId: priTmplId,
		},
		nil,
	)
	return
}

//获取公众号类目
func GetCategory(clt *core.Client, tid string) (err error) {
	response := new(getCategoryResponse)
	_, err = clt.GetJson("/wxaapi/newtmpl/getcategory",
		url.Values{
			"tid": []string{tid}, //模板标题 id，可通过接口获取
		}, response,
	)
	if err != nil {

	}
	return
}

//获取模板中的关键词
func GetPubTemplateKeyWordsById(clt *core.Client, tid string) (err error) {
	response := new(getPubTemplateKeyWordsByIdResponse)
	_, err = clt.GetJson("/wxaapi/newtmpl/getpubtemplatekeywords",
		url.Values{
			"tid": []string{tid}, //模板标题 id，可通过接口获取
		}, response,
	)
	if err != nil {

	}
	return
}

//获取类目下的公共模板
func getPubTemplateTitleList(clt *core.Client, ids []string) (err error) {
	response := new(getPubTemplateTitleListResponse)
	_, err = clt.GetJson("/wxaapi/newtmpl/getpubtemplatetitles",
		url.Values{
			"ids":   []string{}, //类目 id，多个用逗号隔开
			"start": []string{}, //用于分页，表示从 start 开始，从 0 开始计数
			"limit": []string{}, //用于分页，表示拉取 limit 条记录，最大为 30
		}, response,
	)
	if err != nil {

	}
	return
}

//获取私有模板列表
func GetTemplateList(clt *core.Client) (err error) {
	response := new(getTemplateListResponse)
	_, err = clt.GetJson("/wxaapi/newtmpl/gettemplate",
		url.Values{
			"ids":   []string{}, //类目 id，多个用逗号隔开
			"start": []string{}, //用于分页，表示从 start 开始，从 0 开始计数
			"limit": []string{}, //用于分页，表示拉取 limit 条记录，最大为 30
		}, response,
	)
	if err != nil {

	}
	return
}

//发送订阅通知
func send(clt *core.Client, userOpenId string, templateId string) (err error) {
	_, err = clt.PostJson("/cgi-bin/message/subscribe/bizsend",
		nil,
		&sendPostBody{
			ToUser:      userOpenId,
			TemplateId:  templateId,
			Page:        "",
			MiniProgram: nil,
			Data:        nil,
		}, nil,
	)
	return
}
