package base64

import (
	"encoding/base64"
	"testing"
)

func TestEncoding(t *testing.T) {
	msg := "hello,世界"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	t.Log(encoded)
}

func TestDecode(t *testing.T) {
	msg, err := base64.StdEncoding.DecodeString("aGVsbG8s5LiW55WM")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(msg))
}

func TestUrlEncode(t *testing.T) {
	url := "https://www.baidu.com"
	encoded := base64.URLEncoding.EncodeToString([]byte(url))
	t.Log(encoded)
}

func TestUrlDecode(t *testing.T) {
	urlDe := "aHR0cHM6Ly93d3cuYmFpZHUuY29t"
	url, err := base64.URLEncoding.DecodeString(urlDe)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(url))
}
