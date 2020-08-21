package go_test

import (
	"strings"
	"testing"
)

func Case1() {
	str := "1"
	str += "2"
	str += "3"
	str += "4"
	str += "5"
}


func Case2() {
	sbuf := strings.Builder{}
	sbuf.WriteString("1")
	sbuf.WriteString("2")
	sbuf.WriteString("3")
	sbuf.WriteString("4")
	sbuf.WriteString("5")
	sbuf.String()
}

func Test_case1(t *testing.T) {
	Case1()
}

// go test -bench=Benchmark_case1 -test.count=10 | tee ./old.txt
func Benchmark_case1(b *testing.B) {
	for i:=0; i<b.N; i++ {
		Case1()
	}
}


// go test -bench=Benchmark_case2 -test.count=10 | tee ./new.txt
func Benchmark_case2(b *testing.B) {
	for i:=0; i<b.N; i++ {
		Case2()
	}
}

