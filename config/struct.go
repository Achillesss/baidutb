package config

// C config body
type C struct {
	BdussList  []string `toml:"bdussList"`
	ListURL    string   `toml:"listUrl"`    // get forum list
	FidURL     string   `toml:"fidUrl"`     // get forum id
	TbsURL     string   `toml:"tbsUrl"`     // get tbs
	SignURL    string   `toml:"signUrl"`    // sign
	FDetailURL string   `toml:"fDetailUrl"` // get forum detail
}
