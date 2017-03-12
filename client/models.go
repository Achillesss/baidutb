package client

import "github.com/parnurzeal/gorequest"

type agent struct {
	tiebaConf
	tiebaBody
	err error
	req *gorequest.SuperAgent
}

type tiebaConf struct {
	fidURL    string
	tbsURL    string
	signURL   string
	listURL   string
	bdussList []string
}

type tiebaBody struct {
	params map[string]string // bduss, fid, tbs, sign
	kwList map[string]string
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
	Time       int64         `json:"time"` // time when sign
	CTime      int32         `json:"ctime"`
	Logid      int32         `json:"logid"`
}
