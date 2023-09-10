package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int    `sql:"index"`
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=postgres dbname=gwp password=20020610 sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

//func (comment *Comment) Create() (err error) {
//	if comment.PostId == nil {
//		err = errors.New("post not found")
//		return
//	}
//
//	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id", comment.Content, comment.Author, comment.PostId.Id).Scan(&comment.Id)
//	return
//}
//
//func Posts(limit int) (posts []Post, err error) {
//	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
//	if err != nil {
//		return
//	}
//	for rows.Next() {
//		post := Post{}
//		err = rows.Scan(&post.Id, &post.Content, &post.Author)
//		if err != nil {
//			return
//		}
//		posts = append(posts, post)
//	}
//	_ = rows.Close()
//	return
//}
//
//func GetPost(id int) (post Post, err error) {
//	post = Post{}
//	post.Comments = []Comment{}
//	err = Db.QueryRowx("select id, content, author from posts where id = $1", id).StructScan(&post)
//
//	rows, err := Db.Query("select id, content, author from comments")
//	if err != nil {
//		return
//	}
//	for rows.Next() {
//		comment := Comment{PostId: &post}
//		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
//		if err != nil {
//			return
//		}
//		post.Comments = append(post.Comments, comment)
//	}
//	_ = rows.Close()
//	return
//}
//
//func (post *Post) Create() (err error) {
//	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.Author).Scan(&post.Id)
//	return
//}
//
//func (post *Post) Update() (err error) {
//	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
//	return
//}
//
//func (post *Post) Delete() (err error) {
//	_, err = Db.Exec("delete from posts where id = $id", post.Id)
//	return
//}
