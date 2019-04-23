package main

import (
	"fmt"
	"github.com/artslob/gophercises/07-task/cmd"
	"github.com/boltdb/bolt"
	"log"
	"os"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	db, err := bolt.Open(fmt.Sprintf("%s/.todo-list.bolt.db", home), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = db.Close()
	}()
	cmd.Execute(db)
}
