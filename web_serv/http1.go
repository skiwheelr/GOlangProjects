//package main
//
//import (
//	"io"
//	"net/http"
//)
//
//func hello(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "Hello world!")
//}
//
//func main() {
//	http.HandleFunc("/", hello)
//	http.ListenAndServe(":8000", nil)
//}

//https://thenewstack.io/building-a-web-server-in-go/

package main

import (
"io"
"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
io.WriteString(w, "Hello world!")
}

func main() {
mux := http.NewServeMux()
mux.HandleFunc("/", hello)
http.ListenAndServe(":8000", mux)
}
