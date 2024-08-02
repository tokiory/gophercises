package main

import (
	"cyoa/dictionary"
	"cyoa/story"
	"flag"
	"log"
	"net/http"
)

func main() {
	flag.Parse()

	d, err := dictionary.Parse(*dictionary.FlagDictionary)
	if err != nil {
		panic(err)
	}
	st := story.NewStoryHandler(*d)
	log.Fatal(http.ListenAndServe(":8080", st))
}
