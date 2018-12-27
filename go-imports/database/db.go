package database

import "fmt"

var _db string

func InitDb() {
	_db = "initialized"
}

func GetDb() *string {
	fmt.Println("Get DB:", _db)
	return &_db
}
