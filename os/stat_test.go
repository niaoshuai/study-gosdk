package os

import (
	"os"
	"testing"
)

//stat   函数与 lstat 函数的区别： 当一个文件是符号链接时，lstat 函数返回的是该符号链接本身的信息；而 stat 函数返回的是该链接指向文件的信息。

func TestStat(t *testing.T) {
	// 获取文件(夹)的状态信息
	fileInfo, err := os.Stat("/Users/niaoshuai/.m2/")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(fileInfo.ModTime())
}

func TestLstat(t *testing.T) {
	// 获取文件(夹)的状态信息
	fileInfo, err := os.Lstat("/Users/niaoshuai/.m2/")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(fileInfo.ModTime())
}
