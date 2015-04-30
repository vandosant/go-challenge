package main

import (
  "net/http"
  "fmt"
)

func main() {
  http.HandleFunc("/", IndexHandler)
  http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Photo-mosaic Generator")
}
