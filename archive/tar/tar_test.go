package tar

import (
	"archive/tar"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	tmpTar, err := os.Create("tmp.tar.gz")
	if err != nil {
		t.Fatal(err)
	}
	// tar writer
	tw := tar.NewWriter(tmpTar)
	// create file
	// 定义 文件列表
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "please readme"},
		{"readme1.txt", "please readme \n 11"},
		{"readme2.txt", "please readme \n 测试"},
	}

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}

		if err := tw.WriteHeader(hdr); err != nil {
			t.Fatal(err)
		}

		if _, err := tw.Write([]byte(file.Body)); err != nil {
			t.Fatal(err)
		}
	}

	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}

	//clear
	if err := os.Remove("tmp.tar.gz"); err != nil {
		t.Fatal(err)
	}
}

func TestRead(t *testing.T) {
	tarFile, err := os.Open("testdata/tmp.tar.gz")
	if err != nil {
		t.Fatal(err)
	}
	tr := tar.NewReader(tarFile)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("文件名称: %s ", hdr.Name)
		//if _, err := io.Copy(os.Stdout, tr); err != nil {
		//	log.Fatal(err)
		//}
		fd, err := ioutil.ReadAll(tr)
		if err != nil {
			t.Fatal(err)
		}
		// 强制关闭
		t.Log(string(fd))
	}
}
