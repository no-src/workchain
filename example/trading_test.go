package example

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/no-src/workchain"
)

func TestTrading(t *testing.T) {
	buy := workchain.NewWork(buy)
	sell := workchain.NewWork(sell)

	ctx := context.WithValue(context.Background(), "code", "000001")
	ctx = context.WithValue(ctx, "get_current_price", getCurrentPrice)
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		<-time.After(time.Second * 3)
		cancel()
		fmt.Println("stop trading...")
	}()

	workchain.WorkRing(
		buy.WithCond(buyCond(30).Loop()),
		sell.WithCond(sellCond(60).Loop())).
		Do(ctx)

	fmt.Println("trading stopped!")
}

func buy(ctx context.Context) error {
	code := getCode(ctx)
	price := getPrice(ctx)
	fmt.Printf("[%s] [buy] code=%s price=%d\n", now(), code, price)
	return nil
}

func sell(ctx context.Context) error {
	code := getCode(ctx)
	price := getPrice(ctx)
	fmt.Printf("[%s] [sell] code=%s price=%d\n", now(), code, price)
	return nil
}

func buyCond(buyPrice int) workchain.CondFunc {
	return func(ctx context.Context) (bool, context.Context, error) {
		code := getCode(ctx)
		price := getCurrentPriceFunc(ctx)(code)
		if price < buyPrice {
			ctx = setPrice(ctx, price)
			return true, ctx, nil
		}
		return false, ctx, nil
	}
}

func sellCond(sellPrice int) workchain.CondFunc {
	return func(ctx context.Context) (bool, context.Context, error) {
		code := getCode(ctx)
		price := getCurrentPriceFunc(ctx)(code)
		if price > sellPrice {
			ctx = setPrice(ctx, price)
			return true, ctx, nil
		}
		return false, ctx, nil
	}
}

func getCurrentPrice(code string) int {
	time.Sleep(time.Millisecond * 100)
	n := rand.Intn(100)
	if n == 0 {
		return getCurrentPrice(code)
	}
	return n
}

func now() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}

func getCode(ctx context.Context) string {
	return ctx.Value("code").(string)
}

func getPrice(ctx context.Context) int {
	return ctx.Value("price").(int)
}

func setPrice(ctx context.Context, price int) context.Context {
	return context.WithValue(ctx, "price", price)
}

func getCurrentPriceFunc(ctx context.Context) func(code string) int {
	return ctx.Value("get_current_price").(func(code string) int)
}
