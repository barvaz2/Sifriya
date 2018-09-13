package main
import (
  "net/http"
  "strings"
)

func getAllCategoriesList() {
  
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
  if err := http.ListenAndServe(":3000", nil); err != nil {
    panic(err)
  }
}