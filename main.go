package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("views/*.html"))

var temp_posts []Post = make([]Post, 0)

var temp_post_page PostPage = PostPage{
	Posts:  temp_posts,
	Errors: "",
}

func main() {

	for i := 0; i < 10; i++ {
		temp_posts = append(temp_posts, Post{Title: fmt.Sprint("sample Title ", i), Body: fmt.Sprint("sample Body ", i)})
	}

	temp_post_page.Posts = temp_posts

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/posts", posts)
	http.HandleFunc("/about", about)
	log.Fatal(http.ListenAndServe(":80", nil))
}
