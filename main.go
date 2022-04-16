package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/otiai10/opengraph"
)

type handler func(http.ResponseWriter, *http.Request) error

func (f handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	ua := flag.String("user-agent", "Ogjson/1.1", "Value of User-Agent")
	flag.Parse()

	http.Handle("/", handler(func(w http.ResponseWriter, r *http.Request) error {
		url := r.FormValue("url")
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		req.Header.Add("User-Agent", *ua)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		og := opengraph.New(url)
		if err = og.Parse(resp.Body); err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(og); err != nil {
			return err
		}

		return nil
	}))

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
