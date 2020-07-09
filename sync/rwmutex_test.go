package sync

import (
	"sync"
	"testing"
	"time"
)

//https://studygolang.com/articles/2497

//	读写锁是针对读写的互斥锁
//	基本遵循两大原则：
//	1、可以随便读，多个goroutine同时读
//	2、写的时候，啥也不能干。不能读也不能写

var m *sync.RWMutex

func TestReadLock(t *testing.T) {
	m = new(sync.RWMutex)
	// 读
	go read(1, t)
	go read(2, t)
	time.Sleep(2 * time.Second)
}

func TestWriteLock(t *testing.T) {
	m = new(sync.RWMutex)
	// 读
	go write(1, t)
	go read(2, t)
	go write(3, t)
	time.Sleep(10 * time.Second)
}

func read(i int, t *testing.T) {
	t.Log(i, "read start")

	m.RLock()
	t.Log(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()

	t.Log(i, "read over")
}

func write(i int, t *testing.T) {
	t.Log(i, "write start")

	m.Lock()
	t.Log(i, "writing")
	time.Sleep(1 * time.Second)
	m.Unlock()

	t.Log(i, "write over")
}
