package reflect_func

import (
	"fmt"
	"reflect"
	"testing"
)

func Func0(param1 string, param2 int) (int, error) {
	fmt.Println(param1, param2)
	return 1, nil
}


func Func2(arr ...int) int {
	fmt.Println(arr)
	return 0
}


func TestReflectFunc(t *testing.T) {
	var f  = Func0
	inValues := []reflect.Value{reflect.ValueOf("abc"), reflect.ValueOf(11)}
	fv := reflect.ValueOf(f)
	//ft := reflect.TypeOf(f)
	outValues := fv.Call(inValues)
	fmt.Println(outValues, outValues[0].Interface())

	arr := []int{1,2,3,4}
	var f2 = Func2
	fv2 := reflect.ValueOf(f2)
	intValues := []reflect.Value{reflect.ValueOf(arr)}
	outValues = fv2.CallSlice(intValues)
	fmt.Println(outValues)

}


