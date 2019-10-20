package main

import (
	"context"
	"fmt"
	"time"
)

type Key string

func operaionOne(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("context canceled for %s\n", ctx.Value(Key("op_id")))
			return // returning not to leak the goroutine
		default:
			fmt.Printf("OperationOne: %d : opeartion_id = %s\n", n, ctx.Value(Key("op_id")))
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
			fmt.Printf("context canceled for %s\n", ctx.Value(Key("op_id")))
			return // returning not to leak the goroutine
		default:
			fmt.Printf("OperationTwo: %d : opeartion_id = %s\n", n, ctx.Value(Key("op_id")))
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
	ctx = context.WithValue(ctx, Key("op_id"), "ONE")
	go operaionOne(ctx)

	d = time.Now().Add(10000 * time.Millisecond)
	ctx, _ = context.WithDeadline(context.Background(), d)
	ctx = context.WithValue(ctx, Key("op_id"), "TWO")
	go operaionTwo(ctx)

	time.Sleep(20 * time.Second)
}
