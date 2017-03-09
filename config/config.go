package config

import (
	"github.com/BurntSushi/toml"
)

// Decode decode config
func (c *C) Decode(path string) {
	toml.DecodeFile(path, c)
}
