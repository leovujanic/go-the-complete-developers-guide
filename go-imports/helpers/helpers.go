package helpers

import (
	"fmt"
	"github.com/leovujanic/go-the-complete-developers-guide/go-imports/database"
)

func NewHelper() {
	fmt.Println("NewHelper called")
}

func internalHelper() {
	fmt.Println("InternalHelper called")
}

func UseDb() {
	db2 := database.GetDb()
	fmt.Println("DB2", db2)
}
