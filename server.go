package main

import (
  "log"
  "fmt"
  "net/http"
  "encoding/base64"
)

const USER string = "test"
const PASSWORD string = "password"

func handler(w http.ResponseWriter, r *http.Request) {
  authorization := r.Header["Authorization"]
  log.Print(authorization)
  if authorization != nil {
    if checkAuthorization(authorization[0])  {
      fmt.Fprintf(w, "Hello, World")
    } else {
      w.WriteHeader(403)
    }
    
  } else {
    w.Header().Set("WWW-Authenticate", "Basic")
    w.WriteHeader(401)
  }
}
func checkAuthorization(authorization string) bool {
  dec, err := base64.StdEncoding.DecodeString(authorization)
  if err != nil {
    log.Print(err)
    return false
  }
  log.Print(dec)
  return false
}
func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}