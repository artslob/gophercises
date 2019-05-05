package rename

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Renamer interface {
	ContainsPattern(name string) bool
	NewName(name string) string
	Rename(path, oldName, newName string) error
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

func (r DefaultRenamer) Rename(path, oldName, newName string) error {
	return os.Rename(filepath.Join(path, oldName), filepath.Join(path, newName))
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

func (r PrintRenamer) Rename(path, oldName, newName string) error {
	fmt.Printf("in dir %q will rename %q to %q\n", path, oldName, newName)
	return nil
}
