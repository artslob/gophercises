package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/artslob/gophercises/12-renamer/rename"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func flagIsRequired(f *string, name string) {
	if strings.TrimSpace(*f) == "" {
		log.Fatalf("%s flag is required", name)
	}
}

func Walk(root string, renamer rename.Renamer) error {
	if renamer == nil {
		return errors.New("got nil renamer")
	}
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		name := info.Name()
		if !renamer.ContainsPattern(name) {
			return nil
		}
		newName := renamer.NewName(name)
		if err := renamer.Rename(filepath.Dir(path), name, newName); err != nil {
			fmt.Printf("could not rename %q to %q\n", path, newName)
		}
		return nil
	})
}

func main() {
	root := flag.String("root", "", "Root directory to be traversed")
	Old := flag.String("old", "", "String to be renamed in filename")
	New := flag.String("new", "", "What 'old' parameter will be renamed to")
	no := flag.Bool("no", false, "Do not perform any operations that modify the filesystem; print what would happen")
	flag.Parse()
	flagIsRequired(root, "root")
	flagIsRequired(Old, "old")
	flagIsRequired(New, "new")

	var renamer rename.Renamer
	defaultRenamer := rename.NewDefaultRenamer(*Old, *New)
	if *no {
		renamer = rename.NewPrintRenamer(defaultRenamer)
	} else {
		renamer = defaultRenamer
	}
	if err := Walk(*root, renamer); err != nil {
		log.Fatal("got error while walking the path: ", err)
	} else {
		fmt.Println("Done")
	}
}
