package main

import (
  "routes"
  "fmt"
)

func main() {
  a := routes.Compile("/foo/{bar}")
  fmt.Println(a.Matcher)
  fmt.Println(a.Names)

}
