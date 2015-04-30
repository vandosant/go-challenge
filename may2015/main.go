package main

import (
  "net/http"
  "log"
  "fmt"
  "os"
  "io"
)

func main() {
  http.HandleFunc("/files/new", FileCreateHandler)
  http.Handle("/", http.FileServer(http.Dir("public")))
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Photo-mosaic Generator")
}

func FileCreateHandler(w http.ResponseWriter, r *http.Request) {
  r.ParseMultipartForm(32 << 20)
  file, _, err := r.FormFile("file")
  if err != nil {
    fmt.Println(w, err)
    return
  }

  defer file.Close()

  out, err := os.OpenFile("./tmp/testfile", os.O_WRONLY|os.O_CREATE, 0666)
  if err != nil {
    fmt.Println(w, "Unable to create file.")
    return
  }
  defer out.Close()

  _, err = io.Copy(out, file)
  if err != nil {
    fmt.Println(w, err)
    return
  }

  fmt.Println(w, "File uploaded successfully")
}
