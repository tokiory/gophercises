package story

import (
	"cyoa/dictionary"
	"html/template"
	"net/http"
)

var tmpls = template.Must(template.ParseFiles("templates/story.html"))

// isInitialPath checks if the path is the intro path
func getPath(path string) string {
	if path == "/" {
		return "intro"
	}

	return path[1:]
}

type StoryHandler struct {
	dict dictionary.Dictionary
}

func (s *StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	storyName := getPath(r.URL.Path)
	tmpls.ExecuteTemplate(w, "story.html", s.dict[storyName])
}

func NewStoryHandler(dict dictionary.Dictionary) http.Handler {
	return &StoryHandler{
		dict,
	}
}
