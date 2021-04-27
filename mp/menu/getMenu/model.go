package getMenu

type Response struct {
	IsMenuOpen   int64        `json:"is_menu_open"`  //菜单是否开启，0代表未开启，1代表开启
	SelfMenuInfo SelfMenuInfo `json:"selfmenu_info"` //菜单信息
}

type SelfMenuInfo struct {
	Button []Button `json:"button"` //菜单按钮
}

type Button struct {
	Name      string    `json:"name"`           //菜单名称
	Type      string    `json:"type,omitempty"` //菜单的类型，公众平台官网上能够设置的菜单类型有view（跳转网页）、text（返回文本，下同）、img、photo、video、voice。使用API设置的则有8种，详见《自定义菜单创建接口》
	Key       string    `json:"key,omitempty"`  // 菜单KEY值，用于消息接口推送，不超过128字节
	SubButton SubButton `json:"sub_button,omitempty"`
	Value     string    `json:"value,omitempty"`
	URL       string    `json:"url,omitempty"`
	NewsInfo  NewsInfo  `json:"news_info,omitempty"` //图文消息的信息
}

type SubButton struct {
	List []Button `json:"list"`
}

type NewsInfo struct {
	List []Article `json:"list"`
}

type Article struct {
	Title      string `json:"title"`       //图文消息的标题
	Author     string `json:"author"`      //作者
	Digest     string `json:"digest"`      //摘要
	ShowCover  int64  `json:"show_cover"`  //是否显示封面，0为不显示，1为显示
	CoverURL   string `json:"cover_url"`   //封面图片的URL
	ContentURL string `json:"content_url"` //正文的URL
	SourceURL  string `json:"source_url"`  //原文的URL，若置空则无查看原文入口
}
