package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/xxlaefxx/goblog/models"
)

var posts map[string]*models.Post

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"templates/home.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html",
		"templates/leftside.html",
		"templates/rightside.html",
	)
	if err != nil {
		fmt.Println("Template error: ", err)
	}
	t.ExecuteTemplate(w, "home", nil)
}

func newPostHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"templates/post.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html",
		"templates/leftside.html",
		"templates/rightside.html",
	)
	if err != nil {
		fmt.Println("Template error: ", err)
	}
	t.ExecuteTemplate(w, "post", nil)
}

func editPostHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"templates/post.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html",
		"templates/leftside.html",
		"templates/rightside.html",
	)
	if err != nil {
		fmt.Println("Template error: ", err)
	}
	id := r.FormValue("id")
	post, ok := posts[id]
	if !ok {
		http.NotFound(w, r)
	}

	t.ExecuteTemplate(w, "post", post)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"templates/blog.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html",
		"templates/leftside.html",
		"templates/rightside.html",
	)
	if err != nil {
		fmt.Println("Template error: ", err)
	}
	fmt.Println(posts)
	t.ExecuteTemplate(w, "blog", posts)
}

func savePosthandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")
	if id != "" {
		models.EditPost(posts[id], title, content)
	} else {
		post := models.NewPost(generateUUID(), title, content)
		posts[post.ID] = post
	}
	http.Redirect(w, r, "/blog", 302)
}

func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.NotFound(w, r)
	}
	delete(posts, id)
	http.Redirect(w, r, "/blog", 302)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"templates/about.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html",
		"templates/leftside.html",
		"templates/rightside.html",
	)
	if err != nil {
		fmt.Println("Template error: ", err)
	}
	t.ExecuteTemplate(w, "about", nil)
}

func main() {
	fmt.Println("Listening on port 3000")

	posts = make(map[string]*models.Post, 0)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/post", newPostHandler)
	http.HandleFunc("/edit", editPostHandler)
	http.HandleFunc("/delete", deletePostHandler)
	http.HandleFunc("/savepost", savePosthandler)
	http.HandleFunc("/blog", blogHandler)
	http.HandleFunc("/about", aboutHandler)
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics/"))))
	http.ListenAndServe("0.0.0.0:3000", nil)
}
