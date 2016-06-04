//go:generate go-bindata templates/
package main

import (
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	bs "github.com/jubalh/dropf/store/bolt"
)

func main() {
	db, err := bolt.Open("files.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store, _ := bs.NewBoltFileStore(db)

	fh := &FileHandler{}
	fh.Store = store

	// fh := &FileHandler{}
	http.Handle("/", fh)

	// http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/login", loginHandler)
	// http.HandleFunc("/userspace", userspaceHandler)
	// http.HandleFunc("/upload", uploadHandler)

	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
