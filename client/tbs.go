package client

import "encoding/json"

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
	return a.get(a.tbsURL)
}
