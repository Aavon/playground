package utils

import (
	"fmt"
	"io"
	"testing"
	"time"
)

func TestDataPipe_ReadWrite(t *testing.T) {
	pipe := NewDataPipe(10)
	go func() {
		i := 0
		for {
			pipe.Write([]byte(fmt.Sprint(i)))
			i++
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for {
			data, _ := pipe.Read()
			fmt.Println(string(data))
		}
	}()

	time.Sleep(5 * time.Second)
}

func TestIOPipe_ReadWrite(t *testing.T) {
	pipeReader, pipeWriter := io.Pipe()

	go func() {
		i := 0
		for {
			pipeWriter.Write([]byte(fmt.Sprint(i)))
			i++
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for {
			data := make([]byte, 1)
			pipeReader.Read(data)
			fmt.Println(string(data))
		}
	}()

	time.Sleep(5 * time.Second)
}
