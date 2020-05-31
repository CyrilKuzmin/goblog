package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html",
		"templates/leftside.html",
		"templates/rightside.html",
		"templates/home.html",
	)
	if err != nil {
		fmt.Println("Template error: ", err)
	}
	t.ExecuteTemplate(w, "index", nil)
}

func main() {
	fmt.Println("Listening on port 3000")
	http.HandleFunc("/", indexHandler)
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics/"))))
	http.ListenAndServe("0.0.0.0:3000", nil)
}
