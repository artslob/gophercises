package main

import (
	"github.com/artslob/gophercises/12-renamer/rename"
	"testing"
)

type TestRenamer struct {
	rename.DefaultRenamer
	files *[]string
}

func NewTestRenamer(defaultRenamer *rename.DefaultRenamer, files *[]string) *TestRenamer {
	if defaultRenamer == nil {
		return nil
	}
	return &TestRenamer{
		DefaultRenamer: *defaultRenamer,
		files:          files,
	}
}

func (r TestRenamer) Rename(path, oldName, newName string) error {
	*r.files = append(*r.files, newName)
	return nil
}

func TestWalk(t *testing.T) {
	type args struct {
		root    string
		renamer *rename.DefaultRenamer
	}
	tests := []struct {
		name             string
		args             args
		expectedErr      string
		expectedNewNames []string
	}{
		{
			name: "nil renamer",
			args: args{
				root:    "",
				renamer: nil,
			},
			expectedErr: "got nil renamer",
		},
		{
			name: "invalid dir",
			args: args{
				root:    "s",
				renamer: rename.NewDefaultRenamer("day", "pay"),
			},
			expectedErr: "lstat s: no such file or directory",
		},
		{
			name: "sample dir",
			args: args{
				root:    "sample",
				renamer: rename.NewDefaultRenamer("birth", ""),
			},
			expectedNewNames: []string{"day_001.txt", "day_002.txt", "day_003.txt", "day_004.txt"},
		},
		{
			name: "nested",
			args: args{
				root:    "sample",
				renamer: rename.NewDefaultRenamer("n_", "_n"),
			},
			expectedNewNames: []string{"_n008.txt", "_n009.txt", "_n010.txt"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			var files []string
			if tt.args.renamer == nil {
				err = Walk(tt.args.root, nil)
			} else {
				err = Walk(tt.args.root, NewTestRenamer(tt.args.renamer, &files))
			}
			gotError := err != nil
			wantErr := tt.expectedErr != ""
			if gotError != wantErr {
				t.Fatalf("got unexpected error = %q, wantErr %v", err, wantErr)
				return // to disable inspection 'nil pointer dereference' below
			}
			if wantErr && err.Error() != tt.expectedErr {
				t.Fatalf("err not equal to expected: error = %q, expectedErr %q", err, tt.expectedErr)
			}
			if tt.expectedNewNames != nil && !equalFiles(tt.expectedNewNames, files) {
				t.Fatalf("files not equal to expected\nfiles: %q\nexpected: %q", files, tt.expectedNewNames)
			}
		})
	}
}

func equalFiles(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
