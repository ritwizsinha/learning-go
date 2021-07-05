package main

import (
	"fmt"
	"net/http"
	"os"

	"flag"

	"github.com/learning-go/url-shortener/urlshort"
)

func main() {
	mux := defaultMux()
	jsonFilePath := flag.String("json", "", "json")
	yamlFilePath := flag.String("yaml", "", "yaml") 
	flag.Parse()
	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	// Build the YAMLHandler using the mapHandler as the
	// fallback
	fmt.Println(*yamlFilePath, *jsonFilePath)
	if *yamlFilePath != "" {
		yaml, err := os.ReadFile(*yamlFilePath)
		if err != nil {
			fmt.Println(err)
		}
		yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", yamlHandler)
	} else if *jsonFilePath != "" {
		json, err := os.ReadFile(*jsonFilePath)
		if err != nil {
			fmt.Println(err)
		}
		jsonHandler, err := urlshort.JSONHandler(json, mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", jsonHandler)
	} else {
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", mapHandler)
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
