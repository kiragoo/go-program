package gocontext

import (
	"context"
	"fmt"
	"time"
)

func Handler(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handler", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with:", duration)
	}
}
