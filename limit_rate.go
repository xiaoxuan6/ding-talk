package talk

import (
	"sync"
	"time"
)

type LimitRate struct {
	rate  int           `describe:"阀值,从1开始不包含本身"`
	begin time.Time     `describe:"计数开始时间"`
	cycle time.Duration `describe:"计数周期"`
	Count int           `describe:"收到的请求数"`
	lock  sync.Mutex    `describe:"锁"`
}

func NewLimitRate() *LimitRate {
	return new(LimitRate)
}

func (limit *LimitRate) Allow() bool {
	limit.lock.Lock()
	defer limit.lock.Unlock()

	if limit.Count == limit.rate-1 {
		now := time.Now()
		if now.Sub(limit.begin) >= limit.cycle {
			limit.Reset(now)
			return true
		}
		return false
	} else {
		limit.Count++
		return true
	}
}

func (limit *LimitRate) Set(rate int, cycle time.Duration) *LimitRate {
	limit.rate = rate
	limit.begin = time.Now()
	limit.cycle = cycle
	limit.Count = 0
	return limit
}

func (limit *LimitRate) Reset(begin time.Time) {
	limit.begin = begin
	limit.Count = 0
}

func (limit *LimitRate) Reduce() {
	limit.Count = limit.Count - 1
}
