package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	pathPackage "path"
)

func ServeFiles(mux *http.ServeMux, pattern string, addDefaultPaths bool, paths ...string) {
	defaultEnvPaths := []string{"GOPATH", "GOROOT"}
	neededCap := len(paths)
	if addDefaultPaths {
		neededCap += len(defaultEnvPaths)
	}
	if neededCap == 0 {
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
	_, _ = w.Write(content)
}

func (h *SourceFileHandler) readFile(name string) ([]byte, error) {
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