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

	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}
	err = os.WriteFile("create_post.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}

}
