package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
)

func InitDd(driver string, source string, dialect gorp.Dialect) (*gorp.DbMap, error) {
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

	return dbmap, nil
}
