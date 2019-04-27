package main

import (
	"database/sql"
	norm "github.com/artslob/gophercises/08-phone/normalization"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connStr := "postgres://user:pass@localhost?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = db.Close() }()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT id, phone FROM db.phones")
	if err != nil {
		log.Fatal(err)
	}
	phoneSet := make(map[norm.Phone]struct{})
	for rows.Next() {
		var (
			id    int64
			phone string
		)
		if err := rows.Scan(&id, &phone); err != nil {
			log.Fatal(err)
		}
		phoneSet[norm.Normalize(norm.Phone(phone))] = struct{}{}
	}
	_, err = db.Exec("DELETE FROM db.phones")
	if err != nil {
		log.Fatal(err)
	}
	for phone := range phoneSet {
		_, err := db.Exec("INSERT INTO db.phones (phone) VALUES ($1)", phone)
		if err != nil {
			log.Fatal(err)
		}
	}
}
