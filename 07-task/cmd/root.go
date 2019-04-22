package cmd

import (
	"encoding/binary"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var db *bolt.DB

const taskBucket = "taskBucket"

func i64tob(value uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, value)
	return b
}

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Simple cli task manager",
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todo := strings.Join(args, " ")
		err := db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(taskBucket))
			id, _ := b.NextSequence()
			return b.Put(i64tob(id), []byte(todo))
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do!")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(taskBucket))
			if b.Stats().KeyN == 0 {
				fmt.Println("Your TODO list is empty.")
				return nil
			}
			i := 1
			return b.ForEach(func(k, v []byte) error {
				fmt.Printf("%d. %s\n", i, v)
				i++
				return nil
			})
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd, doCmd, listCmd)
}

func Execute(dbParameter *bolt.DB) {
	if dbParameter == nil {
		log.Fatal("db should be non-nil")
	}
	db = dbParameter
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(taskBucket))
		return err
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
