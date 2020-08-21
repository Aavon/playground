package mmap

import (
	"fmt"
	"os"
	"syscall"
	"testing"
)

// unix

func Test_mmap(t *testing.T) {

	var defaultFileSize int64 = 1 << 10

	fd, err := os.OpenFile("./my.mmap", os.O_RDWR|os.O_CREATE, 777)
	defer fd.Close()
	if err != nil {
		panic(err)
	}

	// 指定文件大小
	err = fd.Truncate(defaultFileSize)
	if err != nil {
		panic(err)
	}

	mbuf, err := syscall.Mmap(int(fd.Fd()), 0, int(defaultFileSize), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(mbuf))
	data := []byte("abcdef")
	copy(mbuf, data)
	data[0] = '@'

	fmt.Printf("%T\n", data)

	err = syscall.Munmap(mbuf)
	if err != nil {
		panic(err)
	}
	err = fd.Sync()
	if err != nil {
		panic(err)
	}
}
