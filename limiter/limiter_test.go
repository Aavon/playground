package sync

import (
	"fmt"
	"testing"

	"go.uber.org/ratelimit"
)

func Test_xxx(t *testing.T) {
	var qps int = 600
	var max = 1000
	rate := ratelimit.New(qps)
	ticks := make([]int64, 0)
	for i := 0; i < max; i++ {
		ticks = append(ticks, rate.Take().UnixNano())
	}
	for _, tc := range ticks {
		fmt.Println(tc)
	}
	//data, _ := json.Marshal(ticks)
	//fmt.Println(string(data))
}

// 结论: tick的分布是均匀的
