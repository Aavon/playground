package mapdemo

import (
	"fmt"
	"testing"
)

func TestMapDemo(t *testing.T) {
	var val, r uint32 = 200, 33
	fmt.Println(uint32(uint64(val) * uint64(r) >> 32))
}
