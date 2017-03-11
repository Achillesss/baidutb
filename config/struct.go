package config

// C config body
type C struct {
	Bduss   string `toml:"bduss"`
	ListURL string `toml:"listUrl"`
	FidURL  string `toml:"fidUrl"`
	TbsURL  string `toml:"tbsUrl"`
	SignURL string `toml:"signUrl"`
}
