package client

import (
	"fmt"
	"regexp"

	"github.com/achillesss/baidutb/config"
	"github.com/achillesss/log"
	req "github.com/parnurzeal/gorequest"
)

func parseListResp(resp []byte) (kwList []string) {
	if resp != nil {
		reg := regexp.MustCompile(`\d+\.[<][a]\s\w+[=][[:ascii:]]+[>](\S+|[[:word:]+])[<]\/[a][>]`)
		g := reg.FindAllStringSubmatch(string(resp), -1)
		for i := range g {
			kwList = append(kwList, g[i][1])
		}
	}
	if *debug {
		log.Infofln("kwlist: %s", kwList)
	}
	return
}

func (a *agent) getList(conf *config.C) (res []byte) {
	r := req.New().CustomMethod("GET", conf.ListURL).Set("Cookie", fmt.Sprintf("BDUSS=%s", a.params["bduss"]))
	if *debug {
		r.SetCurlCommand(true)
	}
	_, res, _ = r.EndBytes()
	return
}
