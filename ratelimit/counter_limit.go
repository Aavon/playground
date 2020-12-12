package ratelimit

import (
	"sync"
	"time"
)

type CounterLimiter struct {
	keys        map[string]*token
	intervalSec int
	limit       int
	lock        sync.Mutex
}

type token struct {
	tick      int64
	requested int
}

func NewCounterlimiter(limit, intervalSec int) *CounterLimiter {
	l := &CounterLimiter{
		limit:       limit,
		intervalSec: intervalSec,
	}
	l.keys = make(map[string]*token)
	return l
}

func (l *CounterLimiter) Take(key string) bool {
	if l.limit <= 0 {
		return false
	}
	nowTick := time.Now().Unix() / int64(l.intervalSec)
	l.lock.Lock()
	defer l.lock.Unlock()
	t, ok := l.keys[key]
	if !ok || t.tick != nowTick {
		l.keys[key] = &token{
			tick:      nowTick,
			requested: 1,
		}
		return true
	}

	if t.requested >= l.limit {
		return false
	}
	t.requested += 1
	return true
}
