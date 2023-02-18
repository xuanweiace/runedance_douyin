package TokenBucket

import (
	"github.com/juju/ratelimit"
	"time"
)

/*
ToDo 待压测

*/
import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
)

var limiter = ratelimit.NewBucketWithQuantum(time.Second, 200, 20) //quantum 流速 capacity 容量

func MyTokenBucket() app.HandlerFunc {
	fmt.Println("token create rate:", limiter.Rate())
	fmt.Println("available token :", limiter.Available())
	return func(ctx context.Context, c *app.RequestContext) {
		if limiter.TakeAvailable(1) == 0 {
			log.Printf("available token :%d", limiter.Available())
			c.AbortWithStatusJSON(http.StatusTooManyRequests, "Too Many Request")
		} else {
			c.Next(ctx)
		}
	}
}
