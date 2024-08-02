package urlshort

type HandlerPath struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url" json:"url"`
}

type PathList []HandlerPath

func (pl PathList) Hashmap() map[string]string {
	hashmap := make(map[string]string, len(pl))

	for _, path := range pl {
		hashmap[path.Path] = path.Url
	}

	return hashmap
}
