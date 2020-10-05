package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
					time.Sleep(time.Second)
				}
			}
		}()
		return ch
	}(ctx)

	for v := range ch {
		fmt.Printf("v: %v", v)
		if v == 5 {
			cancel()
			break
		}
	}
}
