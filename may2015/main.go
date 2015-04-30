package main

import (
  "net/http"
  "log"
  "fmt"
)

func main() {
  http.Handle("/", http.FileServer(http.Dir("public")))
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Photo-mosaic Generator")
}
