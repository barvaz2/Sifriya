package main

import (
	"net/http"
	//"io/ioutil"
	"strings"
	"google.golang.org/appengine" // for deploy
)

func getAllCategoriesList() string {
	return "test123"
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	reqpath := strings.TrimPrefix(r.URL.Path, "/")
	sa := strings.Split(reqpath, "/")
	if len(sa) > 1 && sa[0] == "categories" && sa[1] == "list" {
		message := getAllCategoriesList()
		w.Write([]byte(message))
	} else {
		if r.URL.Path != "/" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}
}
func main() {

	// http.HandleFunc("/categories", handleFunc) // for local
	// fs := http.FileServer(http.Dir(""))        // for local
	// http.Handle("/", fs)                       // for local

	http.HandleFunc("/", handleFunc) // for deploy

	// if err := http.ListenAndServe(":3000", nil); err != nil { // for local
	// 	panic(err) // for local
	// } // for local

	appengine.Main() for deploy
}
