package main

import (
	"net/http"
	"net/url"
	"log"
	"fmt"
)

func redirectToHttps(w http.ResponseWriter, req *http.Request) {
	url := &url.URL{
		Scheme: "https",
		Host: req.Host,
		RawPath: req.URL.RawPath,
		RawQuery: req.URL.RawQuery,
		Fragment: req.URL.Fragment,
	}

	log.Println("Redirecting", req.Host)

	http.Redirect(w, req, url.String(), http.StatusMovedPermanently)
}

func health(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	log.Println("Starting…")

	http.HandleFunc("/health", health)
	http.HandleFunc("/", redirectToHttps)

	log.Fatal(http.ListenAndServe(":11111", nil))
}
