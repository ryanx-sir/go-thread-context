package main

import (
	"context"
	"fmt"
	"github.com/ryanx-sir/go-thread-context"
	"time"
)

const ctxKey = "ctxKey"

func main() {
	ctx := context.WithValue(context.Background(), ctxKey, "ctxVal")
	go func() {
		threadContext.SetMeta(ctx, "ctxVal")
		foo()
	}()
	time.Sleep(time.Second)
}

func foo() {
	ctx, data := threadContext.GetMeta()
	fmt.Println("ctxVal:", ctx.Value(ctxKey), "metaData:", data)
}
