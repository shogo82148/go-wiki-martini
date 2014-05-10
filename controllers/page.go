package controllers

import (
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/russross/blackfriday"
	"github.com/shogo82148/go-wiki-martini/models"
	"html/template"
)

func ShowPage(r render.Render, dbmap *gorp.DbMap, params martini.Params) {
	page := &models.Page{}
	dbmap.SelectOne(page, "select * from pages where title = ?", params["title"])
	body := blackfriday.MarkdownBasic([]byte(page.Body))
	r.HTML(200, "show", struct {
		Title string
		Body  template.HTML
	}{
		Title: page.Title,
		Body:  template.HTML(body),
	})
}
