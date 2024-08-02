package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"url-shortener/conf"
	"url-shortener/database"
	"url-shortener/urlshort"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var configFlag = flag.String("c", "", "Path to configuration file with url shortener paths")
var helpFlag = flag.Bool("h", false, "Show help message")

func main() {
	flag.Parse()

	if *helpFlag {
		flag.Usage()
		return
	}

	file, err := os.Open(*configFlag)
	if err != nil {
		panic(err)
	}

	configExtension, err := conf.GetConfExtension(file)
	if err != nil {
		panic(err)
	}

	mux := defaultMux()
	var handler http.HandlerFunc

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	switch configExtension {
	case conf.ConfigExtensionJson:
		content, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		jsonHandler, err := urlshort.JSONHandler(content, mapHandler)
		if err != nil {
			panic(err)
		}
		handler = jsonHandler
	case conf.ConfigExtensionYaml:
		content, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		yamlHandler, err := urlshort.YAMLHandler(content, mapHandler)
		if err != nil {
			panic(err)
		}
		handler = yamlHandler
	case conf.ConfigExtensionDb:
		db, err := database.New(*configFlag)
		if err != nil {
			panic(err)
		}

		dbHandler, err := urlshort.DBHandler(db, mapHandler)
		if err != nil {
			panic(err)
		}
		handler = dbHandler
	}

	fmt.Println("starting the server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
