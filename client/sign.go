package client

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/achillesss/log"
	req "github.com/parnurzeal/gorequest"
)

func (a *agent) signUp(kw string) {
	_, resp, _ := req.New().
		CustomMethod("POST", a.signURL).
		Set("Content-Type", "urlencoded").
		Set("Cookie", fmt.Sprintf("BDUSS=%s", a.params["bduss"])).
		Param("BDUSS", a.params["bduss"]).
		Param("fid", a.params["fid"]).
		Param("kw", kw).
		Param("tbs", a.params["tbs"]).
		Param("sign", a.params["sign"]).
		EndBytes()
	res := new(signRes)
	json.Unmarshal(resp, res)
	desc := "Success"
	if res.ErrMsg != "" {
		desc = res.ErrMsg
	}
	t := time.Unix(res.Time, 0)
	log.Infofln("Sign %q end. Resp: %#v Time: %s", kw, desc, t.Format(time.RFC3339))
}
