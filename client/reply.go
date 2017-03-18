package client

import (
	"fmt"

	"net/http"

	"net/url"
	"strings"

	"io/ioutil"

	"github.com/achillesss/baidutb/config"
	"github.com/achillesss/log"
)

// tbs= 需要
// tid =  帖子id
// fid = 贴吧ID

// curl 'https://tieba.baidu.com/f/commit/post/add' -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' -H 'Cookie: BDUSS=nd2a2h6T2NKRWFvWThKZmNhU3AzSmVSRWJOQmhab1J1fkI2VDhOMG9BdUdrdXRZSVFBQUFBJCQAAAAAAAAAAAEAAACI8wiVX87S0aHW3L3cwtcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIYFxFiGBcRYU;' -H 'Host: tieba.baidu.com' --data 'kw=%E9%BE%99%E7%8F%A0%E6%BF%80%E6%96%97&fid=2396784&tid=4371667081&vcode_md5=&tbs=bd3ba85b331e882c1489836163&content=你只看到我水你的贴，却没看到水贴的原因，你有你的帖子。我有我的回复。你否定我的现在，我决定我的未来。你嘲笑我一无所有，不配回复，我可怜你只会发帖。你可以无视我的回复。但是我会证明这是谁的时代。水贴是注定孤单的旅行。路上少不了鄙夷和不屑，但那又怎样？哪怕不断被删我也要回的漂亮。我是水神。我为自己代言'
func (a *agent) setTid(tid string) *agent {
	return a.setParam("tid", tid)
}
func (a *agent) setContent(content string) *agent {
	return a.setParam("content", content)
}
func (a *agent) reply(conf *config.C) (res []byte) {
	client := new(http.Client)
	body := make(url.Values)
	body.Set("content", a.params["content"])
	body.Add("fid", a.params["fid"])
	body.Add("kw", a.params["kw"])
	body.Add("tid", a.params["tid"])
	body.Add("tbs", a.params["tbs"])
	if *debug {
		log.Infofln("body: %#v\n", body)
	}
	// 写完Header的原因是想看看如何才能让度娘不认为我是机器人
	request, _ := http.NewRequest("POST", conf.ReplyURL, strings.NewReader(body.Encode()))
	request.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	// request.Header.Add("Accept-Encoding", "gzip, deflate, br")
	request.Header.Add("Accept-Language", "en-US,en;q=0.5")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Add("Cookie", fmt.Sprintf("BDUSS=%s", a.params["bduss"]))
	request.Header.Add("Host", "tieba.baidu.com")
	request.Header.Add("Referer", "https://tieba.baidu.com/p/"+a.params["tid"])
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:52.0) Gecko/20100101 Firefox/52.0")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")

	resp, _ := client.Do(request)
	res, _ = ioutil.ReadAll(resp.Body)
	return
}

func replyATopic(bduss, kw, tid string, conf *config.C) {
	a := newAgent().configurate(conf)
	if a.setBduss(bduss).setKw(kw).parseFidResp(a.getFid()).parseTbsResp(a.getTbs()).setTid(tid).setContent(conf.Content).ok() {
		resp := a.reply(conf)
		if *debug {
			log.Infofln("reply response: %s\n", string(resp))
		}
	}
}
