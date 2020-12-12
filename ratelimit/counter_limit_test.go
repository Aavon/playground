package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestCounterLimiter_Limit(t *testing.T) {
	limiter := NewCounterlimiter(3, 3)
	key := "demo"

	for i := 0; i < 5; i++ {
		go func() {
			for {
				if limiter.Take(key) {
					fmt.Println(time.Now().Unix())
				}
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}
	time.Sleep(10 * time.Second)
}
