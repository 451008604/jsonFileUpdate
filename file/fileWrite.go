package file

import (
	"bufio"
	"fmt"
	"os"
)

func FileCreate(path string) *os.File {
	//创建新的文件
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("文件创建失败：", err)
		return nil
	}
	fmt.Println("创建了新的文件：", path)
	return file
}

func FileWrite(path string, str string) {
	// 打开一个文件
	/*
	   const (
	       O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
	       O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
	       O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
	       O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
	       O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	       O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
	       O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
	       O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
	   )
	*/
	// 读写模式打开一个文件并清空内容
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("写入文件时打开错误：", err)
	}

	// 写入时 使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	writer.WriteString(str)
	// 因为writer是带缓存，因此调用WriterString方法时，
	// 内容先写到缓存的，所以调用Flush方法，将缓冲数据
	// 真实写入到文件中，否则文件中没有数据！
	writer.Flush()

	defer file.Close()
}
