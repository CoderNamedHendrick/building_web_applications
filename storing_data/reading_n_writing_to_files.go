package main

import (
	"fmt"
	"os"
)

func runReadingAndWritingToFiles() {
	data := []byte("Hello World!\n")
	err := os.WriteFile("data1", data, 0644)

	if err != nil {
		panic(err)
	}
	read1, _ := os.ReadFile("data1")
	fmt.Println(string(read1))

	file1, _ := os.Create("data2")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}
