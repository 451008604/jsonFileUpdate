package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*文件读取，如果读取不到则会创建对应名称的文件*/
func FileRead(path string) string {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("读取文件时打开错误：", err)
		return ""
	}

	//设置缓存区
	var rb = make([]byte, 1024)

	//读取文件
	n, err := file.Read(rb)
	if err != nil && err != io.EOF {
		fmt.Println("文件读取错误：", err)
		return ""
	}

	//关闭文件
	defer file.Close()

	//文件内容
	fmt.Println("文件读取内容：", string(rb[:n]))
	return string(rb[:n])
}

/*按行读取文件*/
func Bufio(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer file.Close()

	// 创建一个 *Reader ,是带缓冲区的 默认4096
	reader := bufio.NewReader(file)

	// 循环读取文件的内容
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  // io.EOF 表示文件的末尾
			break
		}
		// 输出内容，如果是用Println会多出现一个空行，Println自带换行
		fmt.Printf(str)
	}
	fmt.Println("文件读取结束！")
}

/*检查文件是否存在*/
func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
