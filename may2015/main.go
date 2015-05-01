package main

import (
  "net/http"
  "log"
  "fmt"
  "os"
  "io"
  "encoding/base64"
  "crypto/rand"
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

  id := random(32)


  out, err := os.OpenFile("./tmp/testfile"+id, os.O_WRONLY|os.O_CREATE, 0666)
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

// helpers
func random(size int) string {
  rb := make([]byte,size)
  _, err := rand.Read(rb)


  if err != nil {
     fmt.Println(err)
  }

  rs := base64.URLEncoding.EncodeToString(rb)

  return rs
}
