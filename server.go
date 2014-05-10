package main

import (
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shogo82148/go-wiki-martini/controllers"
	"github.com/shogo82148/go-wiki-martini/models"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	db, _ := models.InitDd("sqlite3", "./db.bin", gorp.SqliteDialect{})
	m.Map(db)

	m.Get("/", func() string {
		return "Hello world!!"
	})
	m.Get("/:title", controllers.ShowPage)

	m.Run()
}
