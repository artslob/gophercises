package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"runtime/debug"
)

type RecoverWrapper struct {
	next http.Handler
	// logStackTrace
	isDev         bool
	lst           bool
	sourcesRegexp *regexp.Regexp
	tmpl          *template.Template
	filesPattern  string
}

func NewRecoverWrapper(next http.Handler, isDev bool, lst bool, filesPattern string) *RecoverWrapper {
	if next == nil {
		panic("got empty handler")
	}
	r := regexp.MustCompile(`((/\S+)+\.go):(\d+)`)
	return &RecoverWrapper{
		next:          next,
		isDev:         isDev,
		lst:           lst,
		sourcesRegexp: r,
		tmpl:          template.Must(template.ParseFiles("recover.html")),
		filesPattern:  filesPattern,
	}
}

func (wr RecoverWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		log.Println(r)
		if wr.lst {
			log.Println(string(debug.Stack()))
		}
		status := http.StatusInternalServerError
		context := struct {
			Status  int
			Err     interface{}
			Message template.HTML
		}{
			Status:  status,
			Err:     r,
			Message: template.HTML(wr.getMessage()),
		}
		var buffer bytes.Buffer
		if err := wr.tmpl.Execute(&buffer, context); err != nil {
			log.Println(err)
			return
		}
		w.WriteHeader(status)
		_, _ = fmt.Fprintln(w, buffer.String())
	}()
	proxyWriter := &MiddlewareResponseWriter{ResponseWriter: w}
	wr.next.ServeHTTP(proxyWriter, r)
	_ = proxyWriter.flush()
}

func (wr RecoverWrapper) getMessage() string {
	if !wr.isDev {
		return "Something went wrong"
	}
	stack := string(debug.Stack())
	stack = wr.sourcesRegexp.ReplaceAllStringFunc(stack, func(s string) string {
		groups := wr.sourcesRegexp.FindStringSubmatch(s)
		query := url.Values{"line": []string{groups[3]}}
		return fmt.Sprintf(`<a href="%s%s?%s">%s:%s</a>`, wr.filesPattern, groups[1], query.Encode(), groups[1], groups[3])
	})
	return fmt.Sprintf("stack:\n%s", stack)
}

type MiddlewareResponseWriter struct {
	http.ResponseWriter
	writes [][]byte
	status int
}

func (w *MiddlewareResponseWriter) Write(block []byte) (int, error) {
	w.writes = append(w.writes, block)
	return len(block), nil
}

func (w *MiddlewareResponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
}

func (w *MiddlewareResponseWriter) flush() error {
	if w.status != 0 {
		w.ResponseWriter.WriteHeader(w.status)
	}
	for _, block := range w.writes {
		if _, err := w.ResponseWriter.Write(block); err != nil {
			return err
		}
	}
	return nil
}
