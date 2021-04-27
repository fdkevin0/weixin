package callback

//CLICK VIEW
type ClickMenuGoToUrlEvent struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
	MenuId       string `xml:"MenuId"`
}

// scancode_push scancode_waitmsg
type ScanCodeEvent struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"`
		ScanResult string `xml:"ScanResult"`
	} `xml:"ScanCodeInfo"`
}

// pic_sysphoto pic_photo_or_album pic_weixin
type pic_sysphoto struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
	SendPicsInfo struct {
		Count   string `xml:"Count"`
		PicList struct {
			Item struct {
				PicMd5Sum string `xml:"PicMd5Sum"`
			} `xml:"item"`
		} `xml:"PicList"`
	} `xml:"SendPicsInfo"`
}

//location_select
type location_select struct {
	ToUserName       string `xml:"ToUserName"`
	FromUserName     string `xml:"FromUserName"`
	CreateTime       string `xml:"CreateTime"`
	MsgType          string `xml:"MsgType"`
	Event            string `xml:"Event"`
	EventKey         string `xml:"EventKey"`
	SendLocationInfo struct {
		LocationX string `xml:"Location_X"`
		LocationY string `xml:"Location_Y"`
		Scale     string `xml:"Scale"`
		Label     string `xml:"Label"`
		Poiname   string `xml:"Poiname"`
	} `xml:"SendLocationInfo"`
}

//view_miniprogram
type view_miniprogram struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
	MenuId       string `xml:"MenuId"`
}
