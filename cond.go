package workchain

import (
	"context"
)

type CondFunc func(ctx context.Context) (bool, context.Context, error)

func (cf CondFunc) Loop() CondFunc {
	return func(ctx context.Context) (bool, context.Context, error) {
		for {
			select {
			case <-ctx.Done():
				return false, ctx, ctx.Err()
			default:
				ok, ctx, err := cf(ctx)
				if err != nil {
					return ok, ctx, err
				}
				if ok {
					return ok, ctx, nil
				}
			}
		}
	}
}
