package config

// C config body
type C struct {
	Bduss        string   `toml:"bduss"`
	TiebaListURL string   `toml:"tiebaListUrl"`
	FidURL       string   `toml:"fidUrl"`
	TbsURL       string   `toml:"tbsUrl"`
	SignURL      string   `toml:"signUrl"`
	FName        []string `toml:"fName"` // tieba name
}
