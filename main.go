package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"coverins/coverage"
	"coverins/logic"
)

type icoverageFlusher interface {
	FlushProfiles()
}

type numberDumper struct {
	value   int
	flusher icoverageFlusher
}

func (n numberDumper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here's your number: %d\n", n.value)
	n.flusher.FlushProfiles()
}

func main() {
	fmt.Println(path.Base(os.Args[0]))
	h := http.NewServeMux()
	f := &coverage.Flusher{}

	h.Handle("/one", numberDumper{1, f})
	h.Handle("/two", numberDumper{2, f})
	h.Handle("/three", numberDumper{3, f})
	h.Handle("/text", logic.Handler{"TEST", f})

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "That's not a supported number!")
		f.FlushProfiles()
	})

	err := http.ListenAndServe(":8080", h)
	log.Fatal(err)
}
