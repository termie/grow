// Package routes provides basic tools for loading dynamic routes from
// a string or file and linking them to http handler functions.

package routes

import (
  "path"
  "regexp"
)

var dynamic_matcher = regexp.MustCompile(`\{([\w_]+)\}`)

type Route struct {
  Matcher *regexp.Regexp
  Names []string
}

func Compile (route string) (*Route) {
  clean_route := path.Clean(route)
  names := dynamic_matcher.FindAllStringSubmatch(clean_route, -1)
  size := len(names)
  names_out := make([]string, size)
  for i, v := range names {
    names_out[i] = v[1]
  }
  repl := dynamic_matcher.ReplaceAllLiteralString(clean_route, `([\w_]+)`)
  return &Route{regexp.MustCompile(repl), names_out}
}
