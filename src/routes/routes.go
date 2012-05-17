// Package routes provides basic tools for loading dynamic routes from
// a string or file and linking them to http handler functions.

package routes

import (
  "path"
  "regexp"
  "net/http"
)

var dynamic_matcher = regexp.MustCompile(`\{([\w_]+)\}`)

type Route struct {
  Handler http.Handler
  Matcher *regexp.Regexp
  Names []string
}


func Compile (url_path string, handler http.Handler) (*Route) {
  clean_path := path.Clean(url_path)

  // Extract the names of the params
  names := dynamic_matcher.FindAllStringSubmatch(clean_path, -1)
  size := len(names)
  names_out := make([]string, size)
  for i, v := range names {
    names_out[i] = v[1]
  }

  // Generate a matcher regexp
  repl := dynamic_matcher.ReplaceAllLiteralString(clean_path, `([\w_]+)`)
  return &Route{handler, regexp.MustCompile(repl), names_out}
}

func (route *Route) Match (url_path string) (bool) {
  return route.Matcher.MatchString(url_path)
}


type Router struct {
  Routes []*Route
}

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  var matched_route *Route

  for _, matcher := range router.Routes {
    if matcher.Match(req.URL.Path) {
      matched_route = matcher
      break
    }
  }
  if matched_route == nil {
    http.Error(w, "No route found for that path.", http.StatusNotFound)
    return
  }
  matched_route.Handler.ServeHTTP(w, req)
}
