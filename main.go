package main

import (
	"html/template"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/xxlaefxx/goblog/db/postsdocuments"
)

var posts *postsdocuments.PostsDocuments

func main() {
	//MongoDB part
	posts = postsdocuments.NewPostsDocuments()

	//Martini (HTTP) part
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
		Directory:       "templates",
		Layout:          "layout",
		Extensions:      []string{".html"},
		Funcs:           []template.FuncMap{unescapeFuncMap},
		Charset:         "UTF-8",
		IndentJSON:      true,
		IndentXML:       true,
		HTMLContentType: "text/html",
	}))
	m.Run()
}
