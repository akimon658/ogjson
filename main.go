package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/otiai10/opengraph"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ogp, err := opengraph.Fetch(r.FormValue("url"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")

		if err = json.NewEncoder(w).Encode(ogp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
