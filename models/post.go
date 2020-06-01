package models

import (
	"time"
)

//Post описывает пост
type Post struct {
	ID         string
	Title      string
	Content    string
	CreateTime string
	ModifyTime string
}

//NewPost создает пост и возвращает указатель на него
func NewPost(id, title, content string) *Post {
	dt := time.Now().Local().Format("01-02-2006 15:04:05")
	return &Post{id, title, content, dt, dt}
}

//EditPost редактируем пост
func EditPost(p *Post, title, content string) {
	p.Title = title
	p.Content = content
	p.ModifyTime = time.Now().Local().Format("01-02-2006 15:04:05")
}
