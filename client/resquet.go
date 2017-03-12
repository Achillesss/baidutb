package client

import (
	"fmt"
	"regexp"
	"time"

	"encoding/json"

	"github.com/achillesss/baidutb/config"
	log "github.com/achillesss/log"
	req "github.com/parnurzeal/gorequest"
)

func (a *agent) doRequest() (resp []byte) {
	if _, resp, _ = a.req.EndBytes(); resp == nil {
		a.err = fmt.Errorf("nil reponse body")
	}
	return
}

func (a *agent) get(url, kw string) *agent {
	if a.err == nil {
		a.req = req.New().CustomMethod("GET", url).Set("Cookie", fmt.Sprintf("BDUSS=%s", a.params["bduss"])).Param("fname", kw)
		if debug {
			s, _ := a.req.AsCurlCommand()
			log.Printfln("[CURL] %s", s)
		}
	}
	return a
}

func (a *agent) parseListResp(resp []byte) *agent {
	if resp != nil {
		reg := regexp.MustCompile(`\d+\.[<][a]\s\w+[=][[:ascii:]]+[>](\S+|[[:word:]+])[<]\/[a][>]`)
		g := reg.FindAllStringSubmatch(string(resp), -1)
		a.kwList = make(map[string]string)
		for i := range g {
			a.kwList[g[i][1]] = time.Now().Format(time.RFC3339)
		}
		log.Printfln("User %s tieba list:\t%s", a.params["bduss"], a.kwList)
	}
	return a
}

func (a *agent) getList() []byte {
	return a.get(a.listURL, "").doRequest()
}

func (a *agent) parseTbsResp(resp []byte) *agent {
	if a.err == nil {
		t := new(tbsRes)
		if a.err = json.Unmarshal(resp, t); a.err == nil {
			a.params["tbs"] = t.Tbs
		}
	}
	return a
}

func (a *agent) getTbs() []byte {
	return a.get(a.tbsURL, "").doRequest()
}

func (a *agent) parseFidResp(resp []byte) *agent {
	if a.err == nil {
		f := new(fidRes)
		if a.err = json.Unmarshal(resp, f); a.err == nil {
			if f.Data != nil {
				a.params["fid"] = fmt.Sprintf("%d", f.Data.Fid)
			} else {
				a.err = fmt.Errorf("nil fid data")
			}
		}
	}
	return a
}
func (a *agent) getFid(kw string) []byte {
	return a.get(a.fidURL, kw).doRequest()
}

func (a *agent) signUp(kw string) time.Time {
	_, resp, _ := req.New().
		CustomMethod("POST", a.signURL).
		Set("Content-Type", "urlencoded").
		Set("Cookie", fmt.Sprintf("BDUSS=%s", a.params["bduss"])).
		Param("BDUSS", a.params["bduss"]).
		Param("fid", a.params["fid"]).
		Param("kw", kw).
		Param("tbs", a.params["tbs"]).
		Param("sign", a.params["sign"]).
		EndBytes()
	res := new(signRes)
	json.Unmarshal(resp, res)
	desc := "Success"
	if res.ErrMsg != "" {
		desc = res.ErrMsg
	}
	t := time.Unix(res.Time, 0)
	log.Printfln("Sign %q end. Resp: %#v Time: %s", kw, desc, t.Format(time.RFC3339))
	return t
}

func (a *agent) startToSign(kw string) {
	if a.err == nil {
		if a.parseFidResp(a.getFid(kw)).ok() {
			signingTimeSlice = append(signingTimeSlice, a.sign(kw).signUp(kw))
		}
	}
}

func transBdussChan(bdussChan chan<- string, bdussList []string) {
	for _, b := range bdussList {
		bdussChan <- b
	}
	close(bdussChan)
}

func transKw(kwChan chan<- string, kwList map[string]string) {
	for k := range kwList {
		kwChan <- k
	}
	close(kwChan)
}

func (a *agent) signByKw(kwChan <-chan string) {
	for k := range kwChan {
		go func(kw string) {
			a.startToSign(kw)
		}(k)
	}
}

func signOnePerson(bduss string, conf *config.C) {
	a := new(agent)
	a.configurate(conf)
	a.params = make(map[string]string)
	a.setBduss(bduss).parseListResp(a.getList()).parseTbsResp(a.getTbs())
	kwChan := make(chan string)
	go transKw(kwChan, a.kwList)
	a.signByKw(kwChan)
	a.log()
}

func signByBDUSS(bdussChan <-chan string, conf *config.C) (res map[string]bool) {
	res = make(map[string]bool)
	for b := range bdussChan {
		go func(bduss string) {
			signOnePerson(bduss, conf)
		}(b)
		res[b] = true
	}
	return
}
