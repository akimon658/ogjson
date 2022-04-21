package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/otiai10/opengraph"
)

type handler func(http.ResponseWriter, *http.Request) *serverError

type serverError struct {
	code    int
	message string
}

func (f handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		http.Error(w, err.message, err.code)
	}
}

func main() {
	ua := flag.String("user-agent", "Ogjson/1.1", "Value of User-Agent")
	flag.Parse()

	http.Handle("/", handler(func(w http.ResponseWriter, r *http.Request) *serverError {
		url := r.FormValue("url")
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return &serverError{http.StatusInternalServerError, err.Error()}
		}
		req.Header.Add("User-Agent", *ua)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return &serverError{http.StatusInternalServerError, err.Error()}
		}

		if resp.StatusCode >= http.StatusBadRequest {
			return &serverError{resp.StatusCode, fmt.Sprintf("error response from %s: %s", url, resp.Status)}
		}

		if !strings.HasPrefix(resp.Header.Get("Content-Type"), "text/html") {
			return &serverError{http.StatusNotFound, `Content type of requested URL is not "text/html"`}
		}

		w.WriteHeader(resp.StatusCode)

		og := opengraph.New(url)
		if err = og.Parse(resp.Body); err != nil {
			return &serverError{http.StatusInternalServerError, err.Error()}
		}

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(og); err != nil {
			return &serverError{http.StatusInternalServerError, err.Error()}
		}

		return nil
	}))

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
