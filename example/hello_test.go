package example

import (
	"context"
	"fmt"
	"testing"

	"github.com/no-src/workchain"
)

func TestHello(t *testing.T) {
	helloWork := workchain.NewWork(hello)
	worldWork := workchain.NewWork(world)
	ctx := context.WithValue(context.Background(), "title", "workchain")
	workchain.WorkChain(helloWork, worldWork).Do(ctx)
}

func hello(ctx context.Context) error {
	fmt.Printf("[%s] hello\n", ctx.Value("title"))
	return nil
}

func world(ctx context.Context) error {
	fmt.Printf("[%s] world\n", ctx.Value("title"))
	return nil
}
