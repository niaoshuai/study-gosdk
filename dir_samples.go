package main

import (
	"fmt"
	"os"
)

func main() {
	// os file
	file, _ := os.Open("/Users/niaoshuai")
	defer file.Close()
	// 读取 文件名称列表
	fileListNames, _ := file.Readdirnames(0)
	fmt.Println(fileListNames)

	file1, _ := os.Open("/Users/niaoshuai")
	defer file1.Close()
	// 读取
	fileInfoList, _ := file1.Readdir(0)
	fmt.Println(fileInfoList)
	// PS: 当 同一个 File Readdirnames 多次调用时候,后面没有数据
}
