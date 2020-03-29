package _defer

import (
	"fmt"
	"testing"
)

//defer:  2
//1
func DeferReturn1() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer: ", i)
	}()
	i++
	return i
}

//defer:  2
//2
func DeferReturn2() (i int) {
	defer func() {
		i++
		fmt.Println("defer: ", i)
	}()
	i++
	return i
}

func TestReturn(t *testing.T) {
	fmt.Println(DeferReturn1())
	fmt.Println(DeferReturn2())
}

// 0
func DeferArg1() {
	var i int
	defer fmt.Println(i)
	i++
}


func DeferArg2() {
	var i int
	defer func() {
		fmt.Println(i)
	}()
	i++
}

func TestArg(t *testing.T) {
	//DeferArg1()
	DeferArg2()
}


func DeferRange1() {
	var arr = []int{1,2,3,4,5}
	for _, i := range arr {
		defer fmt.Println(i)
	}
}

func DeferRange2() {
	var arr = []int{1,2,3,4,5}
	for _, i := range arr {
		defer func(i *int) {
			fmt.Println(*i)
		}(&i)
	}
}

func TestRange(t *testing.T) {
	//DeferRange1()
	DeferRange2()
}




