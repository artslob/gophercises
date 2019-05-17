package main

import (
	"errors"
	"fmt"
	"github.com/alecthomas/chroma/quick"
	"github.com/alecthomas/chroma/styles"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	pathPackage "path"
	"strings"
)

func ServeFiles(mux *http.ServeMux, pattern string, addDefaultPaths bool, paths ...string) {
	defaultEnvPaths := []string{"GOPATH", "GOROOT"}
	neededCap := len(paths)
	if addDefaultPaths {
		neededCap += len(defaultEnvPaths)
	}
	if neededCap == 0 {
		log.Println("serve files did not register any handlers")
		return
	}
	allPaths := make([]string, 0, neededCap)
	allPaths = append(allPaths, paths...)
	if addDefaultPaths {
		for _, goEnv := range defaultEnvPaths {
			if path, ok := os.LookupEnv(goEnv); ok {
				allPaths = append(allPaths, path)
			}
		}
	}
	for _, path := range allPaths {
		prefix := pattern + path + "/"
		mux.Handle(prefix, http.StripPrefix(prefix, &SourceFileHandler{http.Dir(path)}))
	}
}

type SourceFileHandler struct {
	root http.FileSystem
}

func (h *SourceFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := pathPackage.Clean(r.URL.Path)
	log.Print(name)
	content, err := h.readFile(name)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	err = quick.Highlight(w, string(content), "go", "html", styles.Dracula.Name)
	log.Print(err)
}

func (h *SourceFileHandler) readFile(name string) ([]byte, error) {
	if !strings.HasSuffix(name, ".go") {
		return nil, errors.New(fmt.Sprintf("%q is not .go file", name))
	}

	f, err := h.root.Open(name)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if !info.Mode().IsRegular() {
		return nil, errors.New(fmt.Sprintf("%q is not a regular file", name))
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return content, nil
}
