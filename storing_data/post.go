package main

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func storePostToMap(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func storePostsFromFileGob(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func loadPostsFromFileGob(data interface{}, filename string) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func readingAndWritingFiesWithGob() {
	post := Post{Id: 1, Content: "Hello World!", Author: "Sebastine Odeh"}
	storePostsFromFileGob(post, "post1")
	var postRead Post
	loadPostsFromFileGob(&postRead, "post1")
	fmt.Println(postRead)
}

func readingAndWritingToCSV() {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer func(csvFile *os.File) {
		_ = csvFile.Close()
	}(csvFile)

	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "Sebastine Odeh"},
		Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		Post{Id: 3, Content: "Greetings Earthlings!", Author: "Sebastine Odeh"},
		Post{Id: 4, Content: "Hola Mundo!", Author: "Pedro"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}

func runPost() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "Sebastine Odeh"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{Id: 3, Content: "Greetings Earthlings!", Author: "Sebastine Odeh"}
	post4 := Post{Id: 4, Content: "Hola Mundo!", Author: "Pedro"}

	storePostToMap(post1)
	storePostToMap(post2)
	storePostToMap(post3)
	storePostToMap(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Sebastine Odeh"] {
		fmt.Println(post)
	}

	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}
