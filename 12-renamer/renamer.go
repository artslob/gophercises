package main

import (
	"fmt"
	"strings"
)

type Renamer interface {
	ContainsPattern(name string) bool
	NewName(name string) string
	Rename(oldName, newName string)
}

type DefaultRenamer struct {
	old, new string
}

func NewDefaultRenamer(old string, new string) *DefaultRenamer {
	return &DefaultRenamer{old: old, new: new}
}

func (r DefaultRenamer) ContainsPattern(name string) bool {
	return strings.Contains(name, r.old)
}

func (r DefaultRenamer) NewName(name string) string {
	return strings.ReplaceAll(name, r.old, r.new)
}

func (r DefaultRenamer) Rename(oldName, newName string) {
	fmt.Printf("need to implement me!\n")
}

// PrintRenamer overrides Rename method to print new name to stdout
// instead of actually renaming files.
//
// Create like this: NewPrintRenamer(NewDefaultRenamer(*Old, *New))
type PrintRenamer struct {
	DefaultRenamer
}

func NewPrintRenamer(defaultRenamer *DefaultRenamer) *PrintRenamer {
	return &PrintRenamer{DefaultRenamer: *defaultRenamer}
}

func (r PrintRenamer) Rename(oldName, newName string) {
	fmt.Printf("will rename %q to %q\n", oldName, newName)
}
