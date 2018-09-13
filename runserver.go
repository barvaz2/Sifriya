package main
import (
  "net/http"
  "strings"
)
func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message
  w.Write([]byte(message))
}
func main() {

  fs := http.FileServer(http.Dir("public"))
  http.Handle("/", fs)
  http.HandleFunc("/Hello/", sayHello)
  if err := http.ListenAndServe(":3000", nil); err != nil {
    panic(err)
  }
}