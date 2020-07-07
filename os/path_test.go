package os

import (
	"os"
	"testing"
)

func TestMkdirAll(t *testing.T) {
	// 递归创建文件夹
	err := os.MkdirAll("./test/test", os.ModePerm)
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveAll(t *testing.T) {
	// 删除子目录(不包含第一个test)
	err := os.RemoveAll("./test/test")
	if err != nil {
		t.Error(err)
	}
}
