package main
import (
  "net/http"
  "strings"
  "google.golang.org/appengine"
)

func getAllCategoriesList() string {
  return ""
}

func getGategories(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  sa := strings.Split(message, "/")
  if len(sa) > 1 && sa[1] == "list" {
    message = getAllCategoriesList();
    w.Write([]byte(message))
  }
  
}
func main() {

  fs := http.FileServer(http.Dir("public"))
  http.Handle("/", fs)
  http.HandleFunc("/categories/", getGategories)
  if err := http.ListenAndServeTLS(":3000", "ssl_cert/myKey.crt", "ssl_cert/myKey.key", nil); err != nil {
    panic(err)
  }

  appengine.Main()
}