package cache

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestAlign(t *testing.T) {
	type T struct {
		a int32
		//b int16
	}
	var v = struct {
		a []uint16
		b T
	}{}

	fmt.Println(unsafe.Alignof(v), unsafe.Sizeof(v))
}