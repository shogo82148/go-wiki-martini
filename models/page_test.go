package models

import (
	"testing"
)

func TestNewPost(t *testing.T) {
	title := "foobar"
	body := "hogehoge"
	page := NewPage(title, body)
	if page.Title != title {
		t.Errorf("got title %v, want %v", page.Title, title)
	}
	if page.Body != body {
		t.Errorf("got body %v, want %v", page.Body, body)
	}
}

func TestNewPost_dbmap(t *testing.T) {
	title := "foobar"
	body := "hogehoge"
	page := NewPage(title, body)

	dbmap, err := initTestDb()
	if err != nil {
		t.Errorf("db connection error: %v", err)
	}
	dbmap.TruncateTables()
	dbmap.Insert(page)

	page2 := &Page{}
	err = dbmap.SelectOne(page2, "select * from pages where id = ?", page.Id)
	if err != nil {
		t.Errorf("select error: %v", err)
	}
	if page2.Title != title {
		t.Errorf("got title %v, want %v", page.Title, title)
	}
	if page2.Body != body {
		t.Errorf("got body %v, want %v", page.Body, body)
	}
}
