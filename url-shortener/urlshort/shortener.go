package urlshort

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
	"encoding/json"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func (w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
type pathList struct {
	Path string `yaml:"path" json:"path"`
	Url string `yaml:"url" json:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte) ([]pathList, error) {
	var routes []pathList
	if err := yaml.Unmarshal(yml, &routes); err != nil {
		return nil, err
	}
	fmt.Println(routes)
	return routes, nil
}

func buildMap(routes []pathList) map[string]string {
	store := make(map[string]string)
	for _, route := range routes {
		store[route.Path] = route.Url
	}
	return store
}

func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	jsonData, err := parseJSON(json)
	if err != nil {
		return nil ,err
	}
	pathMap := buildMap(jsonData)
	return MapHandler(pathMap, fallback), nil
}

func parseJSON(data []byte) ([]pathList, error) {
	var paths []pathList
	if err := json.Unmarshal(data, &paths); err != nil {
		return nil, err
	}
	return paths, nil
}
