package main

import (
	"context"
	"html/template"
	"net/http"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/xxlaefxx/goblog/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var posts map[string]*models.Post
var client *mongo.Client
var mongoURI = "mongodb://localhost:27017"
var postsCollection *mongo.Collection
var dbName = "blog"
var collectionName = "posts"

func indexHandler(rnd render.Render) {
	rnd.HTML(200, "home", nil)
}

func errorHandler(rnd render.Render) {
	rnd.HTML(400, "error", nil)
}

func blogHandler(rnd render.Render) {
	var posts []models.Post
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	postsCursor, err := postsCollection.Find(ctx, bson.D{})
	if err != nil {
		println("finding error", err)
	}
	for postsCursor.Next(ctx) {
		var elem models.Post
		err := postsCursor.Decode(&elem)
		if err != nil {
			println("decoding error", err)
		}
		posts = append(posts, elem)
	}
	rnd.HTML(200, "blog", posts)
}

func newPostHandler(rnd render.Render) {
	rnd.HTML(200, "post", nil)
}

func aboutHandler(rnd render.Render) {
	rnd.HTML(200, "about", posts)
}

func editPostHandler(rnd render.Render, r *http.Request) {
	id := r.FormValue("id")
	var post models.Post
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	postsCursor, err := postsCollection.Find(ctx, bson.M{"_id": id})
	if err != nil {
		println("finding error", err)
	}
	for postsCursor.Next(ctx) {
		err := postsCursor.Decode(&post)
		if err != nil {
			println("decoding error", err)
		}
	}
	rnd.HTML(200, "post", post)
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
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if id != "" {
		post := models.EditPost(posts[id], title, content, verifyPolicy)
		postsCollection.UpdateOne(ctx, bson.M{"_id": id}, post)
	} else {
		post := models.NewPost(generateUUID(), title, content, verifyPolicy)
		postsCollection.InsertOne(ctx, post)
	}
	rnd.Redirect("/blog", 302)
}

func deletePostHandler(rnd render.Render, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		rnd.Redirect("/error", 302)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	postsCollection.DeleteOne(ctx, bson.M{"_id": id})
	rnd.Redirect("/blog", 302)
}

func main() {
	posts = make(map[string]*models.Post, 0)
	//MongoDB part
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	ctx, mongoConnectCancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer mongoConnectCancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	postsCollection = client.Database(dbName).Collection(collectionName)

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
