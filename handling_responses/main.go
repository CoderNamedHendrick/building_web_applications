package main

import (
	json2 "encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/body", body)
	http.HandleFunc("/process", process)
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}

type Post struct {
	User    string
	Threads []string
}

// sending data in response body
func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

// sending response with status code
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

// sending response tweaking the headers
func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

// returning json response
func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Sebastine",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json2.Marshal(post)
	w.Write(json)
}

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, h)
}

func body(w http.ResponseWriter, r *http.Request) {
	cLen := r.ContentLength
	body := make([]byte, cLen)
	if _, err := r.Body.Read(body); err != nil {
		_, _ = fmt.Fprintln(w, string(body))
		return
	}

	_, _ = fmt.Fprintln(w, "An error occurred")
}

// dealing with multipart forms
func process(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := io.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
	//_ = r.ParseMultipartForm(1024)
	//fileHeader := r.MultipartForm.File["uploaded"][0]
	//file, err := fileHeader.Open()
	//if err == nil {
	//	data, err := io.ReadAll(file)
	//	if err == nil {
	//		_, _ = fmt.Fprintln(w, string(data))
	//	}
	//}
}
