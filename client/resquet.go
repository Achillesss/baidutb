package client

import (
	"encoding/json"

	"fmt"

	"crypto/md5"

	"net/http"

	req "github.com/parnurzeal/gorequest"
)

type fidData struct {
	Fid         int32 `json:"fid"`
	CanSendPics bool  `json:"can_send_pics"`
}
type fidRes struct {
	No   int32    `json:"no"`
	Err  string   `json:"err"`
	Data *fidData `json:"data"`
}
type tbsRes struct {
	Tbs     string `json:"tbs"`
	IsLogin bool   `json:"is_login"`
}

type signErr struct {
	ErrNo   int32  `json:"errno"`
	ErrMsg  string `json:"errmsg"`
	UserMsg string `json:"usermsg"`
}
type signRes struct {
	ErrCode    string                 `json:"error_code"`
	ErrMsg     string                 `json:"error_msg"`
	Error      signErr                `json:"error"`
	Info       map[string]interface{} `json:"info"`
	ServerTime int32                  `json:"server_time"`
	Time       int64                  `json:"time"`
	CTime      int32                  `json:"ctime"`
	Logid      int32                  `json:"logid"`
}

func GetTbs(url, bduss string) (res string) {
	r := req.New()
	r.CustomMethod("GET", url)
	c := http.Cookie{}
	c.Name = "BDUSS"
	c.Value = bduss
	r.AddCookie(&c)
	// r.Header = map[string]string{"Cookie": fmt.Sprintf("BDUSS=%s", bduss)}
	_, data, _ := r.EndBytes()
	t := new(tbsRes)
	json.Unmarshal(data, t)
	res = t.Tbs
	fmt.Printf("TBS RESPONSE: %#v\n", t)
	return
}

// GetFid get tieba id by tieba name
func GetFid(url, kw, bduss string) (res string) {
	r := req.New()
	c := http.Cookie{}
	c.Name = "BDUSS"
	c.Value = bduss
	r.AddCookie(&c)
	_, data, _ := r.Get(url).Param("fname", kw).EndBytes()
	f := new(fidRes)
	json.Unmarshal(data, f)
	if f.Data != nil {
		res = fmt.Sprintf("%d", f.Data.Fid)
	}
	return
}
func Sign(url, bduss, fid, kw, tbs string) (res string) {
	fmt.Printf("bduss:%s\n", bduss)
	fmt.Printf("fid: %s\n", fid)
	fmt.Printf("kw: %s\n", kw)
	fmt.Printf("tbs: %s\n", tbs)
	s := fmt.Sprintf("BDUSS=%sfid=%skw=%stbs=%s", bduss, fid, kw, tbs)
	res = fmt.Sprintf("%X", md5.Sum([]byte(s+"tiebaclient!!!")))
	fmt.Printf("sign: %s\n", res)
	r := req.New()
	r.CustomMethod("POST", url)
	r.Set("Content-Type", "urlencoded")
	r.Param("BDUSS", bduss)
	r.Param("fid", fid)
	r.Param("kw", kw)
	r.Param("tbs", tbs)
	r.Param("sign", res)
	r.Set("Cookie", fmt.Sprintf("BDUSS=%s", bduss))
	r.SetCurlCommand(true)
	curl, _ := r.AsCurlCommand()
	fmt.Printf("CURL: %s\n", curl)
	_, data, _ := r.EndBytes()
	sRes := new(signRes)
	sRes.Info = make(map[string]interface{})
	json.Unmarshal(data, sRes)
	fmt.Printf("response: %#v\n", string(data))
	fmt.Printf("response: %#v\n", sRes)
	return
}
