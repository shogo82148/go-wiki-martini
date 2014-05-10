package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/russross/blackfriday"
	"github.com/shogo82148/go-wiki-martini/models"
	"html/template"
)

func ShowPage(r render.Render, db *models.Database, params martini.Params) {
	page, _ := db.FindPageByTitle(params["title"])
	body := blackfriday.MarkdownBasic([]byte(page.Body))
	r.HTML(200, "show", struct {
		Title string
		Body  template.HTML
	}{
		Title: page.Title,
		Body:  template.HTML(body),
	})
}
