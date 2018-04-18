package logic

import (
	"fmt"
	"net/http"
)

type icoverageFlusher interface {
	FlushProfiles()
}

// Handler ...
type Handler struct {
	Value   string
	Flusher icoverageFlusher
}

func (n Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here's your string: %s\n", n.Value)
	n.Flusher.FlushProfiles()
}
