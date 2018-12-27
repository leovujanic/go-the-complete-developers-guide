package main

import (
	"fmt"
	"github.com/leovujanic/go-the-complete-developers-guide/go-imports/database"
	"github.com/leovujanic/go-the-complete-developers-guide/go-imports/helpers"
)

func main() {

	fmt.Println("Hello")

	database.InitDb()
	db1 := database.GetDb()
	fmt.Println("DB1", db1)

	helpers.NewHelper()

	helpers.UseDb()

}
