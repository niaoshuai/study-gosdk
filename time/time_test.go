package time

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	// 当前时间
	t.Log(time.Now())
	// 解析区间
	h, _ := time.ParseDuration("5h20m")
	t.Logf("hours %.1f", h.Hours())
	// 自定义时间
	t1 := time.Date(2020, time.August, 10, 23, 0, 0, 0, time.UTC)
	t.Log(t1.Local())
	// parse
	t2, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	t.Log(t2)
}
