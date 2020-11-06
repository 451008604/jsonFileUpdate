package main

import (
	"fmt"
	"jsonFileUpdate/file"
	"net/http"
)

func main() {

	//path := "fileFolder/test1.json"
	path := "https://godlike.h5qlyx.com/lyconfig/lyqptest/rule.json"
	//先获取文件原始内容
	file.FileCreate(path)
	file.FileWrite(path, "123")
	file.FileRead(path)

}

// handleFunc
func handleFunc() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "<h1> hello world 1%s</h1>", request.FormValue("name"))
	})

	http.ListenAndServe(":8888", nil)
}
