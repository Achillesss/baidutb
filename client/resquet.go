package client

import (
	"fmt"

	log "github.com/achillesss/log"

	"encoding/json"

	"time"

	req "github.com/parnurzeal/gorequest"
)

func (a *agent) get(url string) *agent {
	if a.err == nil {
		r := req.New().CustomMethod("GET", url).Set("Cookie", fmt.Sprintf("BDUSS=%s", a.Bduss)).Param("fname", a.Kw)
		if a.debug {
			s, _ := r.AsCurlCommand()
			log.Printf("[curl] %s", s)
		}
		_, a.apiResp, _ = r.EndBytes()
	}
	return a
}

func (a *agent) getList() *agent {
	return a.get(a.ListURL).checkResp().log()
}
func (a *agent) getTbs() *agent {
	return a.get(a.tbsURL).checkResp().log()
}

func (a *agent) getFid() *agent {
	return a.get(a.fidURL).checkResp().log()
}

func (a *agent) signUp() time.Time {
	_, a.apiResp, _ = req.New().
		CustomMethod("POST", a.SignURL).
		Set("Content-Type", "urlencoded").
		Set("Cookie", fmt.Sprintf("BDUSS=%s", a.Bduss)).
		Param("BDUSS", a.Bduss).
		Param("fid", a.Fid).
		Param("kw", a.Kw).
		Param("tbs", a.Tbs).
		Param("sign", a.Sign).
		EndBytes()
	res := new(signRes)
	json.Unmarshal(a.apiResp, res)
	desc := "Success"
	if res.ErrMsg != "" {
		desc = res.ErrMsg
	}
	log.Printf("sign %q end.Resp: %#v", a.Kw, desc)
	return time.Unix(res.Time, 0)
}
