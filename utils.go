package main

import (
	"html/template"

	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

//generateUUID генерирует UUID (128 bits)
func generateUUID() string {
	return uuid.New().String()
}

//unescape используется для отображения поста в красивом HTML вместо чистого кода HTML
func unescape(x string) interface{} {
	return template.HTML(x)
}

//MDToHTML принимает на вход строку формата MD и возвращает кусок HTML (строку)
func MDToHTML(md string) string {
	htmlBytes := blackfriday.MarkdownBasic([]byte(md))
	return string(htmlBytes)
}

//MakeNewPolicy создает новую политику bluemonday для верификации пользовательских данных, полученных на метод /savepost
func MakeNewPolicy() *bluemonday.Policy {
	p := bluemonday.UGCPolicy()
	p.AllowAttrs("style").OnElements("span", "p")
	return p
}
