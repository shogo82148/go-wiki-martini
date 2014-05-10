package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
)

type Database struct {
	dbmap *gorp.DbMap
}

func InitDd(driver string, source string, dialect gorp.Dialect) (*Database, error) {
	// create new connection
	db, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: dialect}

	// table settings
	dbmap.AddTableWithName(Page{}, "pages").SetKeys(true, "Id")

	// create tables
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}

	database := &Database{
		dbmap: dbmap,
	}
	return database, nil
}
