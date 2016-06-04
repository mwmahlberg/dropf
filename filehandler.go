package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jubalh/dropf/store"
)

// FileHandler is a simple http.Handler for dealing with file uploads
type FileHandler struct {
	Store store.FileStorer
}

// ServeHTTP displays an upload form and the list of uploaded
// files on a GET request and stores an uploaded for POST requests.
func (fh FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		data, err := Asset("templates/userspace.html")

		if err != nil {
			log.Println("Error loading template: " + err.Error())
		}

		t := template.Must(template.New("upload").Parse(string(data)))
		t.Execute(w, struct{ Files []store.MetaData }{fh.Store.List()})

	case "POST":
		fh.saveFile(r, w)

	}
	return
}

func (fh FileHandler) saveFile(r *http.Request, w http.ResponseWriter) {

	r.ParseMultipartForm(65536)
	f, h, err := r.FormFile("file")

	if err != nil {
		w.WriteHeader(500)
		return
	}

	sf := store.File{}
	sf.Content = f
	sf.Metadata.Name = h.Filename
	sf.Metadata.Public = true
	fh.Store.Store(&sf)
	http.Redirect(w, r, "/", http.StatusFound)
}
