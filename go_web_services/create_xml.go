package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type CreatePost struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

func main() {
	post := CreatePost{
		Id:      "1",
		Content: " Hello World!",
		Author: Author{
			Id:   "2",
			Name: "Sebastine Odeh",
		},
	}

	xmLFile, err := os.Create("create_post.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}
	encoder := xml.NewEncoder(xmLFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding XML to file:", err)
		return
	}

}
