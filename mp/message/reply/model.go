package reply

import "WeChat-golang/mp/core"

type replyTextMessage struct {
	*core.Message        //MsgType text
	Content       string `xml:"Content"` //回复的消息内容（换行：在content中能够换行，微信客户端就支持换行显示）
}

type replyImageMessage struct {
	*core.Message //MsgType image
	Image         struct {
		MediaId string `xml:"MediaId"` //通过素材管理中的接口上传多媒体文件，得到的id
	} `xml:"Image"`
}

type replyVoiceMessage struct {
	*core.Message //MsgType voice
	Voice         struct {
		MediaId string `xml:"MediaId"` //通过素材管理中的接口上传多媒体文件，得到的id
	} `xml:"Voice"`
}

type replyVideoMessage struct {
	*core.Message //MsgType video
	Video         struct {
		MediaId     string `xml:"MediaId"`     //通过素材管理中的接口上传多媒体文件，得到的id
		Title       string `xml:"Title"`       //视频消息的标题
		Description string `xml:"Description"` //视频消息的描述
	} `xml:"Video"`
}

type replyMusicMessage struct {
	*core.Message //MsgType music
	Music         struct {
		Title        string `xml:"Title"`        //音乐标题
		Description  string `xml:"Description"`  //音乐描述
		MusicUrl     string `xml:"MusicUrl"`     //音乐链接
		HQMusicUrl   string `xml:"HQMusicUrl"`   //高质量音乐链接，WIFI环境优先使用该链接播放音乐
		ThumbMediaId string `xml:"ThumbMediaId"` //缩略图的媒体id，通过素材管理中的接口上传多媒体文件，得到的id
	} `xml:"Music"`
}

type replyNewsMessage struct {
	*core.Message     //MsgType news
	ArticleCount  int `xml:"ArticleCount"` //图文消息个数；当用户发送文本、图片、语音、视频、图文、地理位置这六种消息时，开发者只能回复1条图文消息；其余场景最多可回复8条图文消息
	Articles      struct {
		Item []struct {
			Title       string `xml:"Title"`       //图文消息标题
			Description string `xml:"Description"` //图文消息描述
			PicUrl      string `xml:"PicUrl"`      //图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200
			Url         string `xml:"Url"`         //点击图文消息跳转链接
		}
	} `xml:"Articles"` //图文消息信息，注意，如果图文数超过限制，则将只发限制内的条数
}
