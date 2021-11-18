package main

import (
	"fmt"
	"gee/orm"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Age  int
	Name string
}

func main() {
	dsn := "root:shanyu.1028@tcp(49.234.35.185:3306)/gee"
	engine, _ := orm.NewEngine("mysql", dsn)
	defer engine.Close()

	s := engine.NewSession()

	var u user

	result := s.Raw("SELECT * FROM User WHERE name = ?", "Tom").QueryRow()
	result.Scan(&u.Name, &u.Age)

	fmt.Printf("name:%s age:%d\n", u.Name, u.Age)
}
