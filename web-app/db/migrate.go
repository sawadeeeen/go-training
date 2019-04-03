package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sawadeeeen/go-training/web-app/models"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	db.CreateTable(&models.User{})
}
