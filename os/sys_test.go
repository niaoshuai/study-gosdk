package os

import (
	"os"
	"testing"
)

func TestHostname(t *testing.T) {
	// 获取 主机名字
	hostname, err := os.Hostname()
	if err != nil {
		t.Error(err)
	}
	t.Log(hostname)
}
