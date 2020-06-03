package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/xxlaefxx/goblog/models"
)

var posts map[string]*models.Post

func indexHandler(rnd render.Render) {
	rnd.HTML(200, "home", nil)
}

func errorHandler(rnd render.Render) {
	rnd.HTML(400, "error", nil)
}

func blogHandler(rnd render.Render) {
	rnd.HTML(200, "blog", posts)
	for key, value := range posts {
		fmt.Printf("%v : %v\n", key, value)
	}
}

func newPostHandler(rnd render.Render) {
	rnd.HTML(200, "post", nil)
}

func aboutHandler(rnd render.Render) {
	rnd.HTML(200, "about", posts)
}

func editPostHandler(rnd render.Render, r *http.Request) {
	id := r.FormValue("id")
	post, ok := posts[id]
	if !ok {
		rnd.Redirect("/error", 302)
	}
	rnd.HTML(200, "post", post)
}

func savePosthandler(rnd render.Render, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("editor")
	verifyPolicy := MakeNewPolicy()
	if title == "" {
		rnd.Redirect("/error", 302)
		return
	}
	if id != "" {
		models.EditPost(posts[id], title, content, verifyPolicy)
	} else {
		post := models.NewPost(generateUUID(), title, content, verifyPolicy)
		posts[post.ID] = post
	}
	rnd.Redirect("/blog", 302)
}

func deletePostHandler(rnd render.Render, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		rnd.Redirect("/error", 302)
	}
	delete(posts, id)
	rnd.Redirect("/blog", 302)
}

func main() {
	posts = make(map[string]*models.Post, 0)
	unescapeFuncMap := template.FuncMap{"unescape": unescape}
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
	m.Use(render.Renderer(render.Options{
		Directory:       "templates",                         // Specify what path to load the templates from.
		Layout:          "layout",                            // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions:      []string{".html"},                   // Specify extensions to load for templates.
		Funcs:           []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
		Charset:         "UTF-8",                             // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON:      true,                                // Output human readable JSON
		IndentXML:       true,                                // Output human readable XML
		HTMLContentType: "text/html",                         // Output XHTML content type instead of default "text/html"
	}))

	m.Run()
}
