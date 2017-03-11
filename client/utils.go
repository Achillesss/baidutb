package client

import (
	"encoding/json"
	"time"

	"fmt"

	"crypto/md5"
	"encoding/xml"

	"github.com/achillesss/baidutb/config"
	"github.com/achillesss/log"
)

func (a *agent) configurate(c *config.C) *agent {
	a.tiebaConf.ListURL = c.ListURL
	a.tiebaConf.fidURL = c.FidURL
	a.tiebaConf.SignURL = c.SignURL
	a.tiebaConf.tbsURL = c.TbsURL
	a.tiebaBody.Bduss = c.Bduss
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
	case a.Bduss == "":
		a.err = fmt.Errorf("nil bduss")
	}
	return a
}

func (a *agent) checkResp() *agent {
	if a.apiResp == nil {
		a.err = fmt.Errorf("nil reponse body")
	} else {
		if a.debug {
			log.Printf("response: %#v\n", string(a.apiResp))
		}
	}

	return a
}
func (a *agent) parseListResp() *agent {
	l := new(listRes)
	if a.err = xml.Unmarshal(a.apiResp, l); a.err == nil {
		a.apiResp = nil
	}
	log.Printf("list: %#v\n", l)
	return a.log()
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
		case a.Kw == "":
			a.err = fmt.Errorf("nil kw")
		}
	}
	return a.err == nil
}
func (a *agent) sign() *agent {
	a.Sign = fmt.Sprintf("%X", md5.Sum([]byte(fmt.Sprintf("BDUSS=%sfid=%skw=%stbs=%stiebaclient!!!", a.Bduss, a.Fid, a.Kw, a.Tbs))))
	return a
}

func tomorrow(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 500000000, now.Location())
}
