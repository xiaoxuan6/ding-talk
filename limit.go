package talk

import (
	"sync"
	"time"
)

const LimitName = "dinging"

var limit sync.Map

func Init(rate int, cycle time.Duration) {
	limit.Store(LimitName, NewLimitRate().Set(21, 1*time.Minute)) // 官方文档默认限制每分钟20次
}
