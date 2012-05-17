// Package routes provides basic tools for loading dynamic routes from
// a string or file and linking them to http handler functions.

package routes

import (
  "path"
  "regexp"
)

var dynamic_matcher = regexp.MustCompile(`(\{([\w_]+)\})`)


func Compile (route string) (*regexp.Regexp) {
  clean_route := path.Clean(route)
  repl := dynamic_matcher.ReplaceAllLiteralString(clean_route, `([\w_]+)`)
  return regexp.MustCompile(repl)
}
