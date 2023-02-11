package TokenBucket

import (
	"github.com/juju/ratelimit"
	"time"
)

/*
令牌桶是没有优先级的，无法让重要的请求先通过
OP可能因为硬件故障去调整资源, 系统负载也会随着在变化, 如果对服务限流进行缩容和扩容，需要人为手动去修改，运维成本比较大
令牌桶只能对局部服务端的限流, 无法掌控全局资源
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
