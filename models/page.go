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

func (db *Database) InsertPage(page *Page) error {
	return db.dbmap.Insert(page)
}

func (db *Database) FindPageByTitle(title string) (*Page, error) {
	page := &Page{}
	err := db.dbmap.SelectOne(page, "select * from pages where title = ?", title)
	if err != nil {
		return nil, err
	}
	return page, nil
}
