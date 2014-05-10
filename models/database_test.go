package models

import (
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

func initTestDb() (*Database, error) {
	return InitDd("sqlite3", "./tmp_post_db.bin", gorp.SqliteDialect{})
}
