package client

import (
	"encoding/xml"
)

type agent struct {
	tiebaConf
	tiebaBody
	err     error
	apiResp []byte
	debug   bool
}

type tiebaConf struct {
	fidURL    string
	tbsURL    string
	SignURL   string
	ListURL   string
	BdussList []string
	KwList    map[string]string
}
type tiebaBody struct {
	Bduss string `json:"BDUSS"`
	Fid   string `json:"fid"` // tieba id
	Tbs   string `json:"tbs"`
	Sign  string `json:"sign"`
}
type Div struct {
	XMLName xml.Name `xml:"div"`
	A       string   `xml:"a"`
}
type meta struct {
	Name      string `xml:"name,attr"`
	HttpEquiv string `xml:"http_equiv"`
	Content   string `xml:"content,attr"`
}
type head struct {
	Metas []meta `xml:"meta"`
	Style string `xml:"style"`
	Title string `xml:"title"`
}

type a struct {
	Href  string `xml:"href"`
	Value string `xml:"a"`
}

type tr struct {
	Class string `xml:"class,attr"`
}

type tbody struct {
	Tr tr `xml:"tr"`
}
type table struct {
	Class string `xml:"class,attr"`
	Body  tbody  `xml:"tbody"`
}
type div struct {
	Class string `xml:"class,attr"`
	Value table  `xml:"table"`
	A     a      `xml:"a"`
}
type divSpace struct {
	Divs []div `xml:"div"`
}
type body struct {
	DivSpace divSpace `xml:"div"`
}
type listRes struct {
	XMLName xml.Name `xml:"html"`
	Head    head     `xml:"head"`
	Body    body     `xml:"body"`
}
type fidData struct {
	Fid         int32 `json:"fid"`
	CanSendPics int32 `json:"can_send_pics"`
}
type fidRes struct {
	No   int32    `json:"no"`
	Err  string   `json:"err"`
	Data *fidData `json:"data"`
}
type tbsRes struct {
	Tbs     string `json:"tbs"`
	IsLogin int32  `json:"is_login"`
}

type signErr struct {
	ErrNo   int32  `json:"errno"`
	ErrMsg  string `json:"errmsg"`
	UserMsg string `json:"usermsg"`
}
type signRes struct {
	ErrCode    string        `json:"error_code"`
	ErrMsg     string        `json:"error_msg"`
	Error      signErr       `json:"error"`
	Info       []interface{} `json:"info"`
	ServerTime int32         `json:"server_time"`
	Time       int64         `json:"time"`
	CTime      int32         `json:"ctime"`
	Logid      int32         `json:"logid"`
}
