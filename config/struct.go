package config

// C config body
type C struct {
	BdussList []string `toml:"bdussList"`
	ListURL   string   `toml:"listUrl"`
	FidURL    string   `toml:"fidUrl"`
	TbsURL    string   `toml:"tbsUrl"`
	SignURL   string   `toml:"signUrl"`
}
