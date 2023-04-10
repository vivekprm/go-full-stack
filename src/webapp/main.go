package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Naive fileserver implementation
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "" || r.URL.Path == "/" {
			r.URL.Path = "/index.html"
		}
		f, err := os.Open("public" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		defer f.Close()
		io.Copy(w, f)
	})
	// nil uses Default mux which is basically default instance of server.
	http.ListenAndServe(":8080", nil)
}