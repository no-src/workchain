package main

import (
	"context"
	"fmt"

	"github.com/no-src/workchain"
)

func main() {
	helloWork := workchain.NewWork(hello)
	worldWork := workchain.NewWork(world)
	mainWork := workchain.WorkChain(helloWork, worldWork)
	ctx := context.WithValue(context.Background(), "title", "workchain")
	mainWork.Do(ctx)
}

func hello(ctx context.Context) error {
	fmt.Printf("[%s] hello\n", ctx.Value("title"))
	return nil
}

func world(ctx context.Context) error {
	fmt.Printf("[%s] world\n", ctx.Value("title"))
	return nil
}
