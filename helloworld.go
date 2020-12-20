package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello world"))
		_, _ = fmt.Fprintf(writer, "hello,world")
		_, _ = io.WriteString(writer, "hello world")
	})
	_ = http.ListenAndServe(":8000", nil)
}
