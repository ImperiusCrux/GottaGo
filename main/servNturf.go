package main

import (
	"net/http"
)

func calledIfIdUrl(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	if id != "01" {
		id = "1337"
	}
	responseString := "<html><body> Die Id ist: " + id + "</body></html>"
	w.Write([]byte(responseString))
}
