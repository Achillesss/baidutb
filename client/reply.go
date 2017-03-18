package client

import (
	req "github.com/parnurzeal/gorequest"
)

// tbs= 需要
// tid =  帖子id
// fid = 贴吧ID

// curl 'https://tieba.baidu.com/f/commit/post/add' -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' -H 'Cookie: BDUSS=nd2a2h6T2NKRWFvWThKZmNhU3AzSmVSRWJOQmhab1J1fkI2VDhOMG9BdUdrdXRZSVFBQUFBJCQAAAAAAAAAAAEAAACI8wiVX87S0aHW3L3cwtcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIYFxFiGBcRYU;' -H 'Host: tieba.baidu.com' --data 'kw=%E9%BE%99%E7%8F%A0%E6%BF%80%E6%96%97&fid=2396784&tid=4371667081&vcode_md5=&tbs=bd3ba85b331e882c1489836163&content=你只看到我水你的贴，却没看到水贴的原因，你有你的帖子。我有我的回复。你否定我的现在，我决定我的未来。你嘲笑我一无所有，不配回复，我可怜你只会发帖。你可以无视我的回复。但是我会证明这是谁的时代。水贴是注定孤单的旅行。路上少不了鄙夷和不屑，但那又怎样？哪怕不断被删我也要回的漂亮。我是水神。我为自己代言'

func (a *agent) reply() {
    r:=req.New().
}
