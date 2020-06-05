package postsdocuments

import (
	"context"
	"fmt"
	"time"

	"github.com/xxlaefxx/goblog/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//PostsDocuments объект для взаимодействия с MongoDB
type PostsDocuments struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

//NewPostsDocuments создает объект для взаимодействия с MongoDB
func NewPostsDocuments() *PostsDocuments {
	var mongoURI = "mongodb://localhost:27017"
	var dbName = "blog"
	var collectionName = "posts"
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
	db := client.Database(dbName)
	collection := db.Collection(collectionName)
	return &PostsDocuments{client, db, collection}
}

//InsertOne инсертит 1 пост в коллекцию
func (pD PostsDocuments) InsertOne(post *models.Post) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	pD.collection.InsertOne(ctx, post)
}

//UpdateOne инсертит 1 пост в коллекцию
func (pD PostsDocuments) UpdateOne(post *models.Post) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	pD.collection.FindOneAndReplace(ctx, bson.M{"_id": bson.M{"$eq": post.ID}}, post)
}

//DeleteByID удаляет 1 пост в коллекцию
func (pD PostsDocuments) DeleteByID(id string) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	_, err := pD.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		fmt.Printf("Cannot delete post %v, %v\n", id, err)
	}
}

//SelectAll возвращает все посты
func (pD PostsDocuments) SelectAll() *[]models.Post {
	var posts []models.Post
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	postsCursor, err := pD.collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Printf("Cannot find docs: %v\n", err)
	}
	for postsCursor.Next(ctx) {
		var elem models.Post
		err := postsCursor.Decode(&elem)
		if err != nil {
			fmt.Printf("Cannot decode docs: %v\n", err)
		}
		posts = append(posts, elem)
	}
	return &posts
}

//SelectByQuery возвращает документы по данному запросу (query)
func (pD PostsDocuments) SelectByQuery(query bson.M) *[]models.Post {
	var posts []models.Post
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	postsCursor, err := pD.collection.Find(ctx, query)
	if err != nil {
		fmt.Printf("Cannot find docs: %v\n", err)
	}
	for postsCursor.Next(ctx) {
		var elem models.Post
		err := postsCursor.Decode(&elem)
		if err != nil {
			fmt.Printf("Cannot decode docs: %v\n", err)
		}
		posts = append(posts, elem)
	}
	return &posts
}

//SelectByID возвращает один пост по его ID
func (pD PostsDocuments) SelectByID(id string) *models.Post {
	var post models.Post
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	postsCursor, err := pD.collection.Find(ctx, bson.M{"_id": id})
	if err != nil {
		fmt.Printf("Cannot find a doc: %v\n", err)
	}
	for postsCursor.Next(ctx) {
		err := postsCursor.Decode(&post)
		if err != nil {
			fmt.Printf("Cannot decode a doc: %v\n", err)
		}
	}
	return &post
}
