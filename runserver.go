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

func handleFunc(w http.ResponseWriter, r *http.Request) {  
  if r.URL.Path != "/" {
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }

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

  //fs := http.FileServer(http.Dir("public"))
  //http.Handle("/", fs)
  http.HandleFunc("/", handleFunc)
  /*if err := http.ListenAndServeTLS(":443", "ssl_cert/myKey.crt", "ssl_cert/myKey.key", nil); err != nil {
    panic(err)
  }*/

  /*if err := http.ListenAndServe(":3000", nil); err != nil {
    panic(err)
  }*/

  appengine.Main()
}