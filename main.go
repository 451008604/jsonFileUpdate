package main

import (
	"jsonFileUpdate/file"
)

func main() {

	//path := "fileFolder/test1.json"
	path := "https://godlike.h5qlyx.com/lyconfig/lyqptest/rule.json"
	//先获取文件原始内容
	file.FileCreate(path)
	file.FileWrite(path, "123")
	file.FileRead(path)

}
