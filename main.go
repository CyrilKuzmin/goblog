package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-martini/martini"
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

func errorHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"templates/error.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html",
		"templates/leftside.html",
		"templates/rightside.html",
	)
	if err != nil {
		fmt.Println("Template error: ", err)
	}
	t.ExecuteTemplate(w, "error", nil)
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
	if title == "" {
		http.Redirect(w, r, "/error", 302)
		return
	}
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
	posts = make(map[string]*models.Post, 0)

	m := martini.Classic()

	m.Get("/", indexHandler)
	m.Get("/error", errorHandler)
	m.Get("/post", newPostHandler)
	m.Get("/edit", editPostHandler)
	m.Get("/delete", deletePostHandler)
	m.Get("/blog", blogHandler)
	m.Get("/about", aboutHandler)

	m.Post("/savepost", savePosthandler)

	m.Use(martini.Static("statics", martini.StaticOptions{Prefix: "statics"}))

	m.Run()
}
