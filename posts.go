package main

import (
	"net/http"
)

type Post struct {
	Title string
	Body  string
}

type PostPage struct {
	Posts  []Post
	Errors string
}

func posts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := templates.ExecuteTemplate(w, "posts.html", &temp_post_page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodPost {
		requestErrors := PostValidateRequest(r)
		if len(requestErrors) != 0 {
			err := templates.ExecuteTemplate(w, "post_form", requestErrors)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {
			PostCreate(r)
			err := templates.ExecuteTemplate(w, "posts", temp_posts)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	} else {
		http.Error(w, http.ErrNoLocation.Error(), http.StatusBadGateway)
	}
}

func PostValidateRequest(r *http.Request) string {
	title := r.FormValue("title")
	var requestErrors string = ""
	if title == "" {
		requestErrors = requestErrors + "Title Required\n"
	}

	return requestErrors
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
