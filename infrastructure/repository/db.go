package repository

import (
	"gorm.io/gorm"
)

type DB struct {
	Db	*gorm.DB
}

func (db *DB) NewPostgresDB() {
	//database, err := gorm.Open(postgres.Open(""), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	//db.Db = database
}
