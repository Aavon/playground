package sync

import (
	"fmt"
	"go.uber.org/ratelimit"
	"testing"
)

func Test_xxx(t *testing.T) {
	rate := ratelimit.New(1)
	fmt.Println(rate.Take())
	fmt.Println(rate.Take())
	fmt.Println(rate.Take())
	fmt.Println(rate.Take())
	fmt.Println(rate.Take())
	fmt.Println(rate.Take())

}
