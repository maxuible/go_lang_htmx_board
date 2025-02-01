package main

import (
	"net/http"
)

func posts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := templates.ExecuteTemplate(w, "posts.html", temp_posts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodPost {
		PostCreate(r)
		err := templates.ExecuteTemplate(w, "posts", temp_posts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, http.ErrNoLocation.Error(), http.StatusBadGateway)
	}
}

func PostCreate(r *http.Request) {
	post := Post{
		Title: r.FormValue("title"),
		Body:  r.FormValue("body"),
	}
	temp_posts = append([]Post{post}, temp_posts...)
	if len(temp_posts) > 10 {
		temp_posts = append(temp_posts[:10], temp_posts[11:]...)
	}
}
