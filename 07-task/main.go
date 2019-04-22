package main

import (
	"github.com/artslob/gophercises/07-task/cmd"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = db.Close()
	}()
	cmd.Execute(db)
}
