package main

import "fmt"

func main() {
	post := Post{Content: "Hello World!", Author: "Sebastine Odeh"}
	fmt.Println(post)

	Db.Create(&post)
	fmt.Println(post)

	comment := Comment{Content: "Good post!", Author: "Joe"}
	err := Db.Model(&post).Association("Comments").Append(comment).Error
	if err != nil {
		panic(err)
	}

	var readPost Post
	Db.Where("author = $1", "Sebastine Odeh").First(&readPost)
	fmt.Println(readPost)

	var comments []Comment
	Db.Model(&readPost).Related(&comments)
	fmt.Println(comments[0])
}
