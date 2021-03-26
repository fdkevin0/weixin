package standard

import "WeChat-golang/mp/core"

type TextMessage struct {
	*core.StandardMessage        //MsgType text
	Content               string `xml:"Content"` //文本消息内容
}

type ImageMessage struct {
	*core.StandardMessage        //MsgType image
	PicUrl                string `xml:"PicUrl"`  //图片链接（由系统生成）
	MediaId               int64  `xml:"MediaId"` //图片消息媒体id，可以调用获取临时素材接口拉取数据。
}

type VoiceMessage struct {
	*core.StandardMessage        //MsgType voice
	MediaId               int64  `xml:"MediaId"`     //语音消息媒体id，可以调用获取临时素材接口拉取数据。
	Format                string `xml:"Format"`      //语音格式，如amr，speex等
	Recognition           string `xml:"Recognition"` //语音识别结果，UTF8编码
}

type VideoMessage struct {
	*core.StandardMessage        //MsgType video
	MediaId               int64  `xml:"MediaId"`      //视频消息媒体id，可以调用获取临时素材接口拉取数据。
	ThumbMediaId          string `xml:"ThumbMediaId"` //视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
}

type ShortVideoMessage struct {
	*core.StandardMessage        //MsgType shortvideo
	MediaId               int64  `xml:"MediaId"`      //视频消息媒体id，可以调用获取临时素材接口拉取数据。
	ThumbMediaId          string `xml:"ThumbMediaId"` //视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
}

type LocationMessage struct {
	*core.StandardMessage         //MsgType location
	LocationX             float64 `xml:"Location_X"` //地理位置纬度
	LocationY             float64 `xml:"Location_Y"` //地理位置经度
	Scale                 float64 `xml:"Scale"`      //地图缩放大小
	Label                 float64 `xml:"Label"`      //地理位置信息
}

type LinkMessage struct {
	*core.StandardMessage        //MsgType link
	Title                 string `xml:"Title"`       //消息标题
	Description           string `xml:"Description"` //消息描述
	Url                   string `xml:"Url"`         //消息链接
}
