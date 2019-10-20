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

func operaionTwo(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context canceled for op-2")
			return // returning not to leak the goroutine
		default:
			fmt.Printf("OperationTwo: %d\n", n)
			time.Sleep(250 * time.Millisecond)
			n++
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	d := time.Now().Add(5000 * time.Millisecond)
	ctx, _ = context.WithDeadline(context.Background(), d)
	go operaionOne(ctx)

	d = time.Now().Add(10000 * time.Millisecond)
	ctx, _ = context.WithDeadline(context.Background(), d)
	go operaionTwo(ctx)

	time.Sleep(3 * time.Second)
}
