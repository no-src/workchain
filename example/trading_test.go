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

	priceContext := context.WithValue(context.Background(), "code", "000001")
	priceContext = context.WithValue(priceContext, "get_price", getPrice)
	priceContext, cancel := context.WithCancel(priceContext)

	go func() {
		<-time.After(time.Second * 3)
		cancel()
		fmt.Println("stop trading...")
	}()

	workchain.WorkRing(
		buy.WithCond(buyCond(30).Loop()),
		sell.WithCond(sellCond(60).Loop())).
		Do(priceContext)

	fmt.Println("trading stopped!")
}

func buy(ctx context.Context) error {
	code := ctx.Value("code").(string)
	price := ctx.Value("price").(int)
	fmt.Printf("[%s] [buy] code=%s price=%d\n", now(), code, price)
	return nil
}

func sell(ctx context.Context) error {
	code := ctx.Value("code").(string)
	price := ctx.Value("price").(int)
	fmt.Printf("[%s] [sell] code=%s price=%d\n", now(), code, price)
	return nil
}

func buyCond(buyPrice int) workchain.CondFunc {
	return func(ctx context.Context) (bool, context.Context, error) {
		code := ctx.Value("code").(string)
		priceFunc := ctx.Value("get_price").(func(code string) int)
		price := priceFunc(code)
		if price < buyPrice {
			ctx = context.WithValue(ctx, "price", price)
			return true, ctx, nil
		}
		return false, ctx, nil
	}
}

func sellCond(sellPrice int) workchain.CondFunc {
	return func(ctx context.Context) (bool, context.Context, error) {
		code := ctx.Value("code").(string)
		priceFunc := ctx.Value("get_price").(func(code string) int)
		price := priceFunc(code)
		if price > sellPrice {
			ctx = context.WithValue(ctx, "price", price)
			return true, ctx, nil
		}
		return false, ctx, nil
	}
}

func getPrice(code string) int {
	time.Sleep(time.Millisecond * 100)
	n := rand.Intn(100)
	if n == 0 {
		return getPrice(code)
	}
	return n
}

func now() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}
