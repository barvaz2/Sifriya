package main
import (
  "net/http"
  //"io/ioutil"
  "strings"
  "google.golang.org/appengine"
)

func getAllCategoriesList() string {
  return "test123"
}

func getGategories(w http.ResponseWriter, r *http.Request) {  
  reqpath := strings.TrimPrefix(r.URL.Path, "/")
  sa := strings.Split(reqpath, "/")
  if len(sa) > 1 && sa[0] == "categories" && sa[1] == "list" {
    message := getAllCategoriesList();
    w.Write([]byte(message))
  } /*else {
    dat, _ := ioutil.ReadFile(reqpath)
   
      w.Write(dat);
    
    //w.Write([]byte("test123"))
  }*/
}
func main() {

  fs := http.FileServer(http.Dir("public"))
  http.Handle("/", fs)
  http.HandleFunc("/categories", getGategories)
  /*if err := http.ListenAndServeTLS(":443", "ssl_cert/myKey.crt", "ssl_cert/myKey.key", nil); err != nil {
    panic(err)
  }*/

  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }

  appengine.Main()
}