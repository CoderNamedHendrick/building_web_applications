package main

import (
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
	server.ListenAndServe()
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
