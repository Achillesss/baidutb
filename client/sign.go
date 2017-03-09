package client

import (
	"github.com/Achillesss/baidutb/config"
)

func Start(path string) {
	c := new(config.C)
	c.Decode(path)
	for _, f := range c.FName {
		fid := GetFid(c.FidURL, f, c.Bduss)
		tbs := GetTbs(c.TbsURL, c.Bduss)
		Sign(c.SignURL, c.Bduss, fid, f, tbs)
	}
}
