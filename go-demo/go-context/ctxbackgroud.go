package gocontext

import (
	"context"
	"fmt"
	"time"
)

func CtxBackgroud() {
	ctx := context.Background()
	before := time.Now()
	// case 1 -> timeout: preCtx < childCtx -> 父G退出时，触发子G的退出，
	// preCtx, _ := context.WithTimeout(ctx, 100*time.Millisecond)
	// case 2 -> timeout: preCtx > childCtx -> 子G先退出，此时并不影响，当达到父G的退出时间的时候进行退出
	preCtx, _ := context.WithTimeout(ctx, 500*time.Millisecond)
	go func() {
		childCtx, _ := context.WithTimeout(preCtx, 300*time.Millisecond)
		select {
		case <-childCtx.Done():
			after := time.Now()
			fmt.Println("child during: ", after.Sub(before).Milliseconds())
		}
	}()

	select {
	case <-preCtx.Done():
		after := time.Now()
		fmt.Println("pre during: ", after.Sub(before).Milliseconds())
	}
}
