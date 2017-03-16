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

func get(url, bduss, kw string) (res []byte) {
	r := req.New().CustomMethod("GET", url).Set("Cookie", fmt.Sprintf("BDUSS=%s", bduss)).Param("fname", kw)
	_, res, _ = r.EndBytes()
	if debug {
		s, _ := r.AsCurlCommand()
		log.Infofln("[CURL] %s", s)
	}

	return
}

func parseListResp(resp []byte) (kwList []string) {
	if resp != nil {
		reg := regexp.MustCompile(`\d+\.[<][a]\s\w+[=][[:ascii:]]+[>](\S+|[[:word:]+])[<]\/[a][>]`)
		g := reg.FindAllStringSubmatch(string(resp), -1)
		for i := range g {
			kwList = append(kwList, g[i][1])
		}
		// log.Printfln("User %s tieba list:\t%s", a.params["bduss"], kwList)
	}
	return
}

func getList(bduss string, conf *config.C) []byte {
	return get(conf.ListURL, bduss, "")
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
	return get(a.tbsURL, a.params["bduss"], "")
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
	return get(a.fidURL, a.params["bduss"], kw)
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
	log.Infofln("Sign %q end. Resp: %#v Time: %s", kw, desc, t.Format(time.RFC3339))
	return t
}

func transBdussChan(bdussChan chan<- string, bdussList []string) {
	for _, b := range bdussList {
		bdussChan <- b
	}
	close(bdussChan)
}

func transKw(kwChan chan<- string, kwList []string) {
	for _, k := range kwList {
		kwChan <- k
	}
	close(kwChan)
}

func startToSign(bduss, kw string, conf *config.C) {
	a := new(agent)
	a.params = make(map[string]string)
	if a.configurate(conf).setBduss(bduss).parseTbsResp(a.getTbs()).parseFidResp(a.getFid(kw)).ok() {
		signingTimeSlice = append(signingTimeSlice, a.sign(kw).signUp(kw))
	}
	a.log()
}

func signByKw(bduss string, kwChan <-chan string, conf *config.C) {
	for k := range kwChan {
		go func(kw string) {
			startToSign(bduss, kw, conf)
		}(k)
	}
}

func signOnePerson(bduss string, conf *config.C) {
	kwList := parseListResp(getList(bduss, conf))
	kwChan := make(chan string)
	go transKw(kwChan, kwList)
	signByKw(bduss, kwChan, conf)
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
