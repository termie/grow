package main

import (
  "routes"
  "fmt"
  "net/http"
  "flag"
  "log"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
  flag.Parse()
  f := http.HandlerFunc(func (w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello")
  })

  a := routes.Compile("/foo/{bar}", f)
  router := new(routes.Router)
  router.Routes = []*routes.Route{a}

  http.Handle("/", router)
  err := http.ListenAndServe(*addr, nil)
  if err != nil {
      log.Fatal("ListenAndServe:", err)
  }

}
