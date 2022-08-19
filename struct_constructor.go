package main

import "fmt"

type Blog struct {
	Title string
	Content string
}


func createBlog(title, content string) *Blog {
	b := Blog{title, content}
	return &b
}


func main() {
	blog := createBlog("blog1", "content of blog1")
	fmt.Printf("Created Blog[title=%s, content=%s]\n", blog.Title, blog.Content)
}
