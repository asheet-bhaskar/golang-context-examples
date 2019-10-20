package main

import (
	"context"
	"fmt"
	"time"
)

func operaionOne(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context canceled for op-1")
			return // returning not to leak the goroutine
		default:
			fmt.Printf("OperationOne: %d\n", n)
			time.Sleep(500 * time.Millisecond)
			n++
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go operaionOne(ctx)
	time.Sleep(5 * time.Second)
}
