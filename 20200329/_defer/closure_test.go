package _defer

import (
	"fmt"
	"testing"
)

// https://zhuanlan.zhihu.com/p/92634505

func foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Printf("foo1 val = %d\n", *x)
	}
}


func foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo2 val = %d\n", x)
	}
}

func foo3() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fmt.Printf("foo3 val = %d\n", val)
	}
}


func show(v interface{}) {
	fmt.Printf("foo4 val = %v\n", v)
}
func foo4() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go show(val)
	}
}


func foo5() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go func() {
			fmt.Printf("foo5 val = %v\n", val)
		}()
	}
}

var foo6Chan = make(chan int, 10)
func foo6() {
	for val := range foo6Chan {
		go func() {
			fmt.Printf("foo6 val = %d\n", val)
		}()
	}
}

func foo7(x int) []func() {
	var fs []func()
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fs = append(fs, func() {
			fmt.Printf("foo7 val = %d\n", x+val)
		})
	}
	return fs
}

func foo8(x int) (func(), func()) {
	return func() {
		x = x + 1
		fmt.Printf("foo8-1 val = %d\n", x)
	},func() {
		x = x + 1
		fmt.Printf("foo8-2 val = %d\n", x)
	}
}

func TestClosure(t *testing.T) {

	//var i int = 0
	//foo1(&i)()
	//fmt.Println(i)

	//var i int = 0
	//foo2(i)()
	//fmt.Println(i)

	//foo5()

	//foo6Chan <- 1
	//foo6Chan <- 2
	//foo6Chan <- 3
	//close(foo6Chan)
	//foo6()
	//time.Sleep(time.Second)

	f1, f2 := foo8(1)
	f1()
	f2()


}




