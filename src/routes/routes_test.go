package routes

import (
  "testing"
  "net/http"
)

func Test_Compile_and_Match(t *testing.T) {
  path := "/foo/{bar}/boom"
  handler := http.HandlerFunc(func (r http.ResponseWriter, req *http.Request) {
    // pass
  })
  route := Compile(path, handler)
  if (route.Match("/foo")) {
    t.Error("Should not have matched substring of match.")
  }
  if (!route.Match("/foo/blah/boom")) {
    t.Error("Should have matched.")
  }
  params := route.Parse("/foo/blah/boom")
  if (params["bar"] != "blah") {
    t.Error("Did not parse correctly.")
  }
}
