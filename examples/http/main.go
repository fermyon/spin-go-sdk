package main

import (
	"fmt"
	"net/http"

	spinhttp "github.com/fermyon/spin-go-sdk/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("foo", "bar")

		// fmt.Fprintln(w, "== REQUEST ==")
		// fmt.Fprintln(w, "URL:    ", r.URL)
		// fmt.Fprintln(w, "Method: ", r.Method)
		// fmt.Fprintln(w, "Headers:")
		// for k, v := range r.Header {
		// 	fmt.Fprintf(w, "  %q: %q \n", k, v[0])
		// }

		// body, err := io.ReadAll(r.Body)
		// if err != nil {
		// 	fmt.Fprintln(w, "Body Error: ", err)
		// } else {
		// 	fmt.Fprintln(w, "Body:   ", string(body))
		// }

		fmt.Fprintln(w, "== RESPONSE ==")
		fmt.Fprintln(w, "Hello Fermyon!")
	})
}

func main() {}
