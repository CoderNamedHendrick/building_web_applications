package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	post := Post{
		Id:      1,
		Content: "Hello World!",
		Author: Author{
			Id:   2,
			Name: "Sebastine Odeh",
		},
		Comments: []Comment{
			{
				Id:      3,
				Content: "Have a great day!",
				Author:  "Adam",
			},
			{
				Id:      4,
				Content: "How are you today?",
				Author:  "Betty",
			},
		},
	}

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	err = os.WriteFile("create_post.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

}
