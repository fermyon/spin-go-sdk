package main

import (
	"fmt"
	"net/http"

	spinhttp "github.com/fermyon/spin-go-sdk/v2/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("foo", "bar")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "== RESPONSE ==")
		fmt.Fprintln(w, "Hello Fermyon!")
		fmt.Fprintln(w, "Hello again Fermyon!")
	})
}

func main() {}
