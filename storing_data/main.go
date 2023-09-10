package main

import "fmt"

func main() {
	post := Post{Content: "Hello World!", AuthorName: "Sebastine Odeh"}
	_ = post.Create()

	comment := Comment{Content: "Good post!", Author: "Joe", Post: &post}
	_ = comment.Create()

	readPost, _ := GetPost(post.Id)

	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)
}
