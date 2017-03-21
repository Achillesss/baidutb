package client

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/achillesss/baidutb/config"
	"github.com/achillesss/log"
)

func (a *agent) configurate(c *config.C) *agent {
	a.tiebaConf.listURL = c.ListURL
	a.tiebaConf.fidURL = c.FidURL
	a.tiebaConf.signURL = c.SignURL
	a.tiebaConf.tbsURL = c.TbsURL
	a.tiebaConf.bdussList = c.BdussList
	a.tiebaConf.fDetailURL = c.FDetailURL
	return a
}

func (a *agent) log() *agent {
	log.FmtErrN(1, &a.err)
	log.Lln(a.err != nil, "%#v", a.err)
	return a
}

func (a *agent) checkConf() *agent {
	switch {
	case a == nil:
		a.err = fmt.Errorf("nil agent")
	case a.listURL == "":
		a.err = fmt.Errorf("nil list url")
	case a.fidURL == "":
		a.err = fmt.Errorf("nil fid url")
	case a.tbsURL == "":
		a.err = fmt.Errorf("nil tbs url")
	case a.signURL == "":
		a.err = fmt.Errorf("nil sign url")
	case a.bdussList == nil:
		a.err = fmt.Errorf("nil bduss list")
	}
	return a
}

func (a *agent) ok() bool {
	if a.params["bduss"] != "" && a.params["fid"] != "" && a.params["tbs"] != "" {
		a.err = nil
	} else {
		switch {
		case a.params["bduss"] == "":
			a.err = fmt.Errorf("nil bduss")
		case a.params["fid"] == "":
			a.err = fmt.Errorf("nil fid")
		case a.params["tbs"] == "":
			a.err = fmt.Errorf("nil tbs")
		}
	}
	return a.err == nil
}

func (a *agent) sign(kw string) *agent {
	a.params["sign"] = fmt.Sprintf("%X", md5.Sum([]byte(fmt.Sprintf("BDUSS=%sfid=%skw=%stbs=%stiebaclient!!!", a.params["bduss"], a.params["fid"], kw, a.params["tbs"]))))
	return a
}

func tomorrow(now time.Time) time.Time {
	return today(now).AddDate(0, 0, 1).Add(time.Millisecond * 500)
}
func today(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func countDown() {
	time.Sleep(time.Second)
	for i := 10; i > 0; i-- {
		second := "seconds"
		if i == 1 {
			second = "second"
		}
		log.Infofln("Signing will end in %d %s...", i, second)
		time.Sleep(time.Second)
	}
}
