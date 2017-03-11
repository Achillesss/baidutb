package client

import (
	"fmt"
	"time"

	"encoding/json"

	log "github.com/achillesss/log"
	req "github.com/parnurzeal/gorequest"
)

func (a *agent) get(url, kw string) *agent {
	if a.err == nil {
		r := req.New().CustomMethod("GET", url).Set("Cookie", fmt.Sprintf("BDUSS=%s", a.Bduss)).Param("fname", kw)
		if a.debug {
			s, _ := r.AsCurlCommand()
			log.Printfln("[curl] %s", s)
		}
		_, a.apiResp, _ = r.EndBytes()
	}
	return a
}

func (a *agent) getList() *agent {
	return a.get(a.ListURL, "").checkResp().log()
}
func (a *agent) getTbs(kw string) *agent {
	return a.get(a.tbsURL, kw).checkResp().log()
}

func (a *agent) getFid(kw string) *agent {
	return a.get(a.fidURL, kw).checkResp().log()
}

func (a *agent) signUp(kw string) time.Time {
	_, a.apiResp, _ = req.New().
		CustomMethod("POST", a.SignURL).
		Set("Content-Type", "urlencoded").
		Set("Cookie", fmt.Sprintf("BDUSS=%s", a.Bduss)).
		Param("BDUSS", a.Bduss).
		Param("fid", a.Fid).
		Param("kw", kw).
		Param("tbs", a.Tbs).
		Param("sign", a.Sign).
		EndBytes()
	res := new(signRes)
	json.Unmarshal(a.apiResp, res)
	desc := "Success"
	if res.ErrMsg != "" {
		desc = res.ErrMsg
	}
	t := time.Unix(res.Time, 0)
	log.Printfln("sign %q end.Resp: %#v Time: %s", kw, desc, t.Format(time.RFC3339))
	return t
}

func (a *agent) signOneTieba(kw string) {
	if a.err == nil {
		if a.getFid(kw).parseFidResp().getTbs(kw).parseTbsResp().canSign() {
			now := a.sign(kw).signUp(kw)
			timeChan <- &now
		}
	}
}

func transBdussChan(bdussChan chan<- string, bdussList []string) {
	for _, b := range bdussList {
		bdussChan <- b
	}
	close(bdussChan)
}

func transKwChan(kwChan chan<- string, kwList map[string]string) {
	for k := range kwList {
		kwChan <- k
	}
	close(kwChan)
}

func (a *agent) signOnePerson(bduss string) {
	a.setBduss(bduss)
	a.getList().parseListResp()
	kwChan := make(chan string)
	go transKwChan(kwChan, a.KwList)
	for k := range kwChan {
		go func(kw string) {
			a.signOneTieba(kw)
		}(k)
	}
}
