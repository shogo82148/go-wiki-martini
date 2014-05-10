package models

type Page struct {
	Id    int
	Title string
	Body  string
}

func NewPage(title, body string) *Page {
	return &Page{
		Title: title,
		Body:  body,
	}
}
