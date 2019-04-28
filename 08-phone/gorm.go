package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func mainGorm() {
	db, err := gorm.Open("postgres", "postgres://user:pass@localhost?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = db.Close() }()
}
