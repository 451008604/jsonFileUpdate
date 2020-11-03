package main

import "jsonFileUpdate/file"

func main() {

	path := "fileFolder/test1.json"
	//先获取文件原始内容
	str := document.FileRead(path)

	str += "testWirte  "
	document.FileWrite(path, str)

}
