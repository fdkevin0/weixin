package subMessage

func errorMap(errorCode int) string {
	return map[int]string{
		200014: "模版 tid 参数错误",
		200020: "关键词列表 kidList 参数错误",
		200021: "场景描述 sceneDesc 参数错误",
		200011: "此账号已被封禁，无法操作",
		200013: "此模版已被封禁，无法选用",
		200012: "私有模板数已达上限，上限 50 个",
	}[errorCode]
}
