package controllers

import (
	"../models"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func ShowPage(r render.Render, dbmap *gorp.DbMap, params martini.Params) {
	page := &models.Page{}
	dbmap.SelectOne(page, "select * from pages where title = ?", params["title"])
	r.HTML(200, "show", page)
}
