package main

import (
	"database/sql"
	norm "github.com/artslob/gophercises/08-phone/normalization"
	_ "github.com/lib/pq"
	"log"
)

func mainSql() {
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
	defer func() { _ = rows.Close() }()
	phoneSet := make(map[norm.Phone]struct{})
	for rows.Next() {
		var (
			id    int64
			phone string
		)
		if err := rows.Scan(&id, &phone); err != nil {
			log.Fatal(err)
		}
		normalizedPhone := norm.Phone(phone).Normalize()
		if _, ok := phoneSet[normalizedPhone]; ok {
			if _, err := db.Exec("DELETE FROM db.phones WHERE id = $1", id); err != nil {
				log.Fatal(err)
			}
		} else {
			phoneSet[normalizedPhone] = struct{}{}
			if _, err := db.Exec("UPDATE db.phones SET phone = $1 WHERE id = $2", normalizedPhone, id); err != nil {
				log.Fatal(err)
			}
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
