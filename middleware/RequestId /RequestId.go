package RequestId

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"runedance_douyin/pkg/tools"
)

func MyRequestid() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		tools.LoggerInit()
		requestId := uuid.New()

		c.Set("request_id", requestId.String())

	}
}
