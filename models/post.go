package models

import (
	"time"

	"github.com/microcosm-cc/bluemonday"
)

//Post описывает пост
type Post struct {
	ID          string `bson:"_id,omitempty"`
	Title       string
	ContentHTML string
	CreateTime  string
	ModifyTime  string
}

//NewPost создает пост и возвращает указатель на него
func NewPost(id, title, content string, policy *bluemonday.Policy) *Post {
	println("create new post ", content)
	dt := time.Now().Local().Format("01-02-2006 15:04:05")
	return &Post{id, title, policy.Sanitize(content), dt, dt}
}

//EditPost редактируем пост
func EditPost(p *Post, title, content string, policy *bluemonday.Policy) *Post {
	p.Title = title
	p.ContentHTML = policy.Sanitize(content)
	p.ModifyTime = time.Now().Local().Format("01-02-2006 15:04:05")
	return p
}
