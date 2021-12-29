package gocontext

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// 这里的 Handler 处理时间是大于主线程的过期时间的，所以有机会进行处理业务逻辑代码
	go Handler(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}
