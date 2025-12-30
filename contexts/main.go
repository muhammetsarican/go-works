package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Contexts")

	ch := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		time.Sleep(5 * time.Second)
		ch <- "20 pieces of washing machine"
		close(ch)
	}()

	select {
	case detail := <-ch:
		fmt.Printf("Product details received: %s", detail)
	case <-ctx.Done():
		fmt.Println("Timeout")
	}

	orderedFunctionsCtx := context.Background()
	orderedFunctionsCtx = context.WithValue(orderedFunctionsCtx, "correlationID", "abc1231")

	f1(orderedFunctionsCtx)
}

func f1(ctx context.Context) {
	fmt.Println("f1", ctx.Value("correlationID"))
	f2(ctx)
}
func f2(ctx context.Context) {
	fmt.Println("f2", ctx.Value("correlationID"))
	f3(ctx)
}
func f3(ctx context.Context) {
	fmt.Println("f3", ctx.Value("correlationID"))
}
