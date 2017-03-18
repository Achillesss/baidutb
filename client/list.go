package client

import (
	"regexp"

	"github.com/achillesss/baidutb/config"
)

func parseListResp(resp []byte) (kwList []string) {
	if resp != nil {
		reg := regexp.MustCompile(`\d+\.[<][a]\s\w+[=][[:ascii:]]+[>](\S+|[[:word:]+])[<]\/[a][>]`)
		g := reg.FindAllStringSubmatch(string(resp), -1)
		for i := range g {
			kwList = append(kwList, g[i][1])
		}
	}
	return
}

func (a *agent) getList(conf *config.C) []byte {
	return a.get(conf.ListURL)
}
