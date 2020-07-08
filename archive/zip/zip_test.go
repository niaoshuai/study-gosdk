package zip

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"testing"
)

// 初始化
func TestMain(m *testing.M) {
	// 开始执行测试
	println("测试开始")
	m.Run()
}

func TestAll(t *testing.T) {
	// 1. 创建测试 zip 文件
	// 2. 读取测试 zip 文件
	// 3. 删除测试 zip 文件
}

func TestWriteZip(t *testing.T) {
	zipFile, err := os.Create("tmp.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer zipFile.Close()
	// 创建zip writer
	w := zip.NewWriter(zipFile)
	// 定义 文件列表
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "please readme"},
		{"readme1.txt", "please readme \n 11"},
		{"readme2.txt", "please readme \n 测试"},
	}
	// 遍历添加文件
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			t.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			t.Fatal(err)
		}
	}

	// 必须要检查并关闭 (写入)
	err = w.Close()
	if err != nil {
		t.Fatal(err)
	}

	// 清除测试 zip 文件
	err = os.Remove("tmp.zip")
	if err != nil {
		t.Fatal(err)
	}

}

func TestReadZip(t *testing.T) {
	// 读取 zip
	r, err := zip.OpenReader("testdata/tmp.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	// 输出文件内容
	for _, f := range r.File {
		t.Logf("文件名称是:  %s \n", f.Name)
		rc, err := f.Open()
		if err != nil {
			t.Fatal(err)
		}
		fd, err := ioutil.ReadAll(rc)
		if err != nil {
			t.Fatal(err)
		}
		// 强制关闭
		rc.Close()
		t.Log(string(fd))
	}
}
