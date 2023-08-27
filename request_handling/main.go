package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/http2"
	"net/http"
	"reflect"
	"runtime"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!\n", p.ByName("name"))
}

// chaining requests
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(writer, request)
	}
}

func main() {
	// using mux
	//mux := httprouter.New()
	//mux.GET("/hello/:name", hello)
	//server := http.Server{
	//	Addr:    "127.0.0.1:8080",
	//	Handler: mux,
	//}

	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}

	_ = http2.ConfigureServer(&server, &http2.Server{})
	_ = server.ListenAndServeTLS("cert.pem", "key.pem")
	//_ = server.ListenAndServe()
	//_ = server.ListenAndServeTLS("cert.pem", "key.pem")

}
