package main

import (
	"fmt"
	"net/http"

	"github.com/pressly/chi"
)

func init() {
	calcMem("chi", initChi)
}

func initChi() {
	h := chi.NewRouter()
	h.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("Hello, World"))
	})
	h.Get("/:name", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(fmt.Sprintf("Hello, %s", chi.URLParam(r, "name"))))
	})
	registerHandler("chi", h)
}
