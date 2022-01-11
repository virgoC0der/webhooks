package io

type FeiShuReq struct {
	MsgType string       `json:"msg_type"`
	Content *ContentInfo `json:"content"`
}

type ContentInfo struct {
	Post *PostInfo `json:"post"`
}

type PostInfo struct {
	ZhCn *TextInfo `json:"zh_cn"`
}

type TextInfo struct {
	Title   string        `json:"title"`
	Content []interface{} `json:"content"`
}

type FeiShuResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type A struct {
	Tag      string `json:"tag"`
	UnEscape bool   `json:"un_escape"`
	Href     string `json:"href"`
}

type Text struct {
	Tag      string `json:"tag"`
	Text     string `json:"text"`
	UnEscape bool   `json:"un_escape"`
}

type AT struct {
	UserId string `json:"user_id"`
}

type Img struct {
	ImageKey string `json:"image_key"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}
