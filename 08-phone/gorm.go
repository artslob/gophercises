package main

import (
	norm "github.com/artslob/gophercises/08-phone/normalization"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type PhoneRow struct {
	Id    int        `gorm:"primary_key;auto_increment;not null"`
	Phone norm.Phone `gorm:"type:varchar(30);not null"`
}

func (PhoneRow) TableName() string {
	return "db.phones"
}

func mainGorm() {
	db, err := gorm.Open("postgres", "postgres://user:pass@localhost?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = db.Close() }()
	var rows []PhoneRow
	db.Find(&rows)
	phoneSet := make(map[norm.Phone]struct{})
	for _, row := range rows {
		normalizedPhone := row.Phone.Normalize()
		if _, ok := phoneSet[normalizedPhone]; ok {
			db.Delete(&row)
		} else {
			phoneSet[normalizedPhone] = struct{}{}
			db.Save(&PhoneRow{row.Id, normalizedPhone})
		}
	}
}
