package client

import (
	"encoding/json"
	"fmt"
)

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

func (a *agent) getFid() []byte {
	return a.get(a.fidURL)
}
