package cache

import (
	"sync/atomic"
	"testing"
)

// pad 结构的 x y z 会被并发的执行原子操作
type nopad struct {
	x uint64 // 8byte
	y uint64 // 8byte
	z uint64 // 8byte
}

func (s *nopad) increase() {
	atomic.AddUint64(&s.x, 1)
	atomic.AddUint64(&s.y, 1)
	atomic.AddUint64(&s.z, 1)
}


type pad struct {
	x uint64 // 8byte
	_ [56]byte
	y uint64 // 8byte
	_ [56]byte
	z uint64 // 8byte
	_ [56]byte
}

func (s *pad) increase() {
	atomic.AddUint64(&s.x, 1)
	atomic.AddUint64(&s.y, 1)
	atomic.AddUint64(&s.z, 1)
}

func BenchmarkPad(b *testing.B) {
	//benchstat
	//s := nopad{}
	s := pad{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.increase()
		}
	})
}

func BenchmarkNoPad(b *testing.B) {
	s := nopad{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.increase()
		}
	})
}