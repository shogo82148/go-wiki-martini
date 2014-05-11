package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/russross/blackfriday"
	"github.com/shogo82148/go-wiki-martini/models"
	"html/template"
	"net/http"
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

func CreatePage(r render.Render, db *models.Database, params martini.Params) {
	r.HTML(200, "create", nil)
}

func SaveNewPage(res http.ResponseWriter, req *http.Request, db *models.Database) {
	title := req.FormValue("title")
	body := req.FormValue("body")
	page := models.NewPage(title, body)
	db.InsertPage(page)
	http.Redirect(res, req, "/pages/"+title, http.StatusFound)
}
