package dictionary

type DictionaryOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type DictionaryItem struct {
	Title   string             `json:"title"`
	Options []DictionaryOption `json:"options"`
	Story   []string           `json:"story"`
}

type Dictionary map[string]DictionaryItem
