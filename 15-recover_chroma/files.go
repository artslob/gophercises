package main

import (
	"net/http"
	"os"
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
		mux.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(path))))
	}
}
