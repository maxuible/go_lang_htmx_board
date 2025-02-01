package main

import "net/http"

func about(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./views/about.html")
}
