package config

// C config body
type C struct {
	BdussList  []string `toml:"bdussList"`
	ListURL    string   `toml:"listUrl"`    // get forum list
	FidURL     string   `toml:"fidUrl"`     // get forum id
	TbsURL     string   `toml:"tbsUrl"`     // get tbs
	SignURL    string   `toml:"signUrl"`    // sign
	FDetailURL string   `toml:"fDetailUrl"` // get forum detail
	ReplyURL   string   `toml:"replyUrl"`   // url for posting a reply
	Content    string   `toml:"content"`    // what you want to leave in a topic
}
