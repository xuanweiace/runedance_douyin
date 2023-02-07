package TokenBucket

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"sync"
	"time"
)

type limiter struct {
	limitspeed float64
	burst      float64

	mu       sync.Mutex
	token    float64
	lasttime time.Time
}

func newLimiter(limit float64, burst float64) *limiter {
	return &limiter{limitspeed: limit, burst: burst}
}
func (lim *limiter) Allow() bool {
	return lim.AllowN(time.Now(), 1)
}

func (lim *limiter) AllowN(now time.Time, n int) bool {
	lim.mu.Lock()
	defer lim.mu.Unlock()
	delta := now.Sub(lim.lasttime).Seconds() * float64(lim.limitspeed)
	lim.token += delta
	if lim.token > lim.burst {
		lim.token = lim.burst
	}
	if float64(n) > lim.token {
		return false
	}
	lim.token -= float64(n)
	lim.lasttime = now
	return true
}

var limit = newLimiter(100, 100)

func MyTokenBucket() app.HandlerFunc {
	fmt.Println("token create speed rate: ", limit.limitspeed)
	fmt.Println("token available token", limit.burst)
	return func(ctx context.Context, c *app.RequestContext) {
		if !limit.Allow() {
			log.Println("available token: %d", limit.token)
		} else {
			//c.Set()
		}
	}
}
