package main

import (
	"fmt"
	"net/http"

	"github.com/martini-contrib/render"
	"github.com/xxlaefxx/goblog/models"
)

func indexHandler(rnd render.Render) {
	rnd.HTML(200, "home", nil)
}

func errorHandler(rnd render.Render) {
	rnd.HTML(400, "error", nil)
}

func blogHandler(rnd render.Render) {
	rnd.HTML(200, "blog", posts.SelectAll())
}

func newPostHandler(rnd render.Render) {
	rnd.HTML(200, "post", nil)
}

func aboutHandler(rnd render.Render) {
	rnd.HTML(200, "about", posts)
}

func editPostHandler(rnd render.Render, r *http.Request) {
	id := r.FormValue("id")
	rnd.HTML(200, "post", posts.SelectByID(id))
}

func savePosthandler(rnd render.Render, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("editor")
	if title == "" {
		//Без заголовка не принимаем
		rnd.Redirect("/error", 302)
		return
	}
	verifyPolicy := MakeNewPolicy()
	if id != "" {
		fmt.Println(posts.SelectByID(id))
		post := models.EditPost(posts.SelectByID(id), title, content, verifyPolicy)
		fmt.Println(post)
		posts.UpdateOne(post)
	} else {
		post := models.NewPost(generateUUID(), title, content, verifyPolicy)
		posts.InsertOne(post)
	}
	rnd.Redirect("/blog", 302)
}

func deletePostHandler(rnd render.Render, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		rnd.Redirect("/error", 302)
	}
	posts.DeleteByID(id)
	rnd.Redirect("/blog", 302)
}
