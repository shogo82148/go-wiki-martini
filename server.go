package main

import (
	"./controllers"
	"./models"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	dbmap, _ := models.InitDd("sqlite3", "./db.bin", gorp.SqliteDialect{})
	m.Map(dbmap)

	m.Get("/", func() string {
		return "Hello world!!"
	})
	m.Get("/:title", controllers.ShowPage)

	m.Run()
}
