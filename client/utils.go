package client

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/achillesss/baidutb/config"
	"github.com/achillesss/log"
)

func (a *agent) setBduss(bduss string) *agent {
	a.tiebaBody.Bduss = bduss
	return a
}
func (a *agent) configurate(c *config.C) *agent {
	a.tiebaConf.ListURL = c.ListURL
	a.tiebaConf.fidURL = c.FidURL
	a.tiebaConf.SignURL = c.SignURL
	a.tiebaConf.tbsURL = c.TbsURL
	a.tiebaConf.BdussList = c.BdussList
	return a
}

func (a *agent) log() *agent {
	log.Logn(a.err == nil, fmt.Sprintf("%v", log.FErrorN(a.err, 1)), 1)
	return a
}

func (a *agent) checkConf() *agent {
	switch {
	case a == nil:
		a.err = fmt.Errorf("nil agent")
	case a.ListURL == "":
		a.err = fmt.Errorf("nil list url")
	case a.fidURL == "":
		a.err = fmt.Errorf("nil fid url")
	case a.tbsURL == "":
		a.err = fmt.Errorf("nil tbs url")
	case a.SignURL == "":
		a.err = fmt.Errorf("nil sign url")
	case a.BdussList == nil:
		a.err = fmt.Errorf("nil bduss list")
	}
	return a
}

func (a *agent) checkResp() *agent {
	if a.apiResp == nil {
		a.err = fmt.Errorf("nil reponse body")
	} else {
		if a.debug {
			log.Printfln("response: %#v", string(a.apiResp))
		}
	}

	return a
}

func (a *agent) parseListResp() *agent {
	r := regexp.MustCompile(`\d+\.[<][a]\s\w+[=][[:ascii:]]+[>](\S+|[[:word:]+])[<]\/[a][>]`)
	g := r.FindAllStringSubmatch(string(a.apiResp), -1)
	a.KwList = make(map[string]string)
	for i := range g {
		a.KwList[g[i][1]] = time.Now().Format(time.RFC3339)
	}
	a.apiResp = nil
	log.Printfln("user %s tieba list:\t%s", a.Bduss, a.KwList)
	return a
}
func (a *agent) parseTbsResp() *agent {
	if a.err == nil {
		t := new(tbsRes)
		if a.err = json.Unmarshal(a.apiResp, t); a.err == nil {
			a.Tbs = t.Tbs
			a.apiResp = nil
		}
	}
	return a.log()
}
func (a *agent) parseFidResp() *agent {
	if a.err == nil {
		f := new(fidRes)
		if a.err = json.Unmarshal(a.apiResp, f); a.err == nil {
			if f.Data != nil {
				a.Fid = fmt.Sprintf("%d", f.Data.Fid)
				a.apiResp = nil
			} else {
				a.err = fmt.Errorf("nil fid data")
			}
		}
	}
	return a.log()
}

func (a *agent) canSign() bool {
	if a.err == nil {
		switch {
		case a.Bduss == "":
			a.err = fmt.Errorf("nil bduss")
		case a.Fid == "":
			a.err = fmt.Errorf("nil fid")
		case a.Tbs == "":
			a.err = fmt.Errorf("nil tbs")
			// case a.Kw == "":
			// a.err = fmt.Errorf("nil kw")
		}
	}
	return a.err == nil
}
func (a *agent) sign(kw string) *agent {
	a.Sign = fmt.Sprintf("%X", md5.Sum([]byte(fmt.Sprintf("BDUSS=%sfid=%skw=%stbs=%stiebaclient!!!", a.Bduss, a.Fid, kw, a.Tbs))))
	return a
}

func tomorrow(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 500000000, now.Location())
}
