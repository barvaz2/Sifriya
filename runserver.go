package main
import (
  "net/http"
  //"strings"
  "google.golang.org/appengine"
)

func getAllCategoriesList() string {
  return "test123"
}

func getGategories(w http.ResponseWriter, r *http.Request) {
  /*message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  sa := strings.Split(message, "/")
  if len(sa) > 1 && sa[0] == "categories" && sa[1] == "list" {
    message = getAllCategoriesList();
    w.Write([]byte(message))
  }*/
  w.Write([]byte("test123"))
}
func main() {

  //fs := http.FileServer(http.Dir("public"))
  //http.Handle("/", fs)
  http.HandleFunc("/", getGategories)
  /*if err := http.ListenAndServeTLS(":443", "ssl_cert/myKey.crt", "ssl_cert/myKey.key", nil); err != nil {
    panic(err)
  }*/

  appengine.Main()
}