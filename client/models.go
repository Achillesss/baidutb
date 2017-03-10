package client

type agent struct {
	tiebaConf
	tiebaBody
	err     error
	apiResp []byte
	debug   bool
}

type tiebaConf struct {
	fidURL  string
	tbsURL  string
	SignURL string
}
type tiebaBody struct {
	Bduss string `json:"BDUSS"`
	Fid   string `json:"fid"` // tieba id
	Kw    string `json:"kw"`  // tieba name
	Tbs   string `json:"tbs"`
	Sign  string `json:"sign"`
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
