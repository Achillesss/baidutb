package client

import (
	"fmt"

	"time"

	"strings"

	"github.com/achillesss/baidutb/config"
	log "github.com/achillesss/log"
	req "github.com/parnurzeal/gorequest"
)

// func (a *agent) setBduss(bduss string) *agent {
// 	a.bduss = bduss
// 	return a
// }
func newAgent() *agent {
	a := new(agent)
	a.params = make(map[string]string)
	return a
}

func (a *agent) setParam(key, value string) *agent {
	a.params[key] = value
	return a
}

func (a *agent) setBduss(bduss string) *agent {
	return a.setParam("bduss", bduss)
}

func (a *agent) setKw(kw string) *agent {
	return a.setParam("kw", kw)
}

func (a *agent) get(url string) (res []byte) {
	r := req.New().CustomMethod("GET", url).Set("Cookie", fmt.Sprintf("BDUSS=%s", a.params["bduss"])).Param("fname", a.params["kw"]).Param("kw", a.params["kw"])
	if *debug {
		r.SetCurlCommand(true)
	}
	_, res, _ = r.EndBytes()
	return
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
	a := newAgent()
	if a.configurate(conf).setBduss(bduss).setKw(kw).parseTbsResp(a.getTbs()).parseFidResp(a.getFid()).ok() {
		a.sign(kw).signUp(kw)
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
	a := newAgent().setBduss(bduss)
	kwList := parseListResp(a.getList(conf))
	log.Infofln("User %s tieba list:\t%s", bduss, kwList)
	kwChan := make(chan string)
	go transKw(kwChan, kwList)
	signByKw(bduss, kwChan, conf)
}

func signByBDUSS(bdussChan <-chan string, conf *config.C) (res map[string]bool) {
	res = make(map[string]bool)
	for b := range bdussChan {
		signOnePerson(b, conf)
		res[b] = true
	}
	return
}

func getTopicList(bduss string, conf *config.C) {
	a := newAgent().setBduss(bduss)
	kwList := parseListResp(a.getList(conf))
	// 弱智,ufo
	for _, kw := range kwList {
		go func(fName string) {
			b := newAgent().configurate(conf).setBduss(bduss).setKw(fName)
			topicList := parseTopicListResp(b.getTopicList())
			// log.Infofln("帖子列表：%#v", topicList)
		topic:
			for k, v := range topicList {
				// 水一个贴，帖子名包含"吧规"的不水
				unReply := []string{"吧规", "禁止水", "精品", "置顶"}
				for _, u := range unReply {
					if strings.Contains(v, u) {
						continue topic
					}
				}
				resp := replyATopic(bduss, fName, k, conf)
				log.Infofln("用户%q在贴吧%q水贴：%s\t结果：%#v", bduss[:10], fName, v, resp)
				time.Sleep(2 * time.Second)
			}
		}(kw)
	}
}
