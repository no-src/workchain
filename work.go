package workchain

import "context"

type Work struct {
	prev *Work
	next *Work

	do   DoFunc
	cond CondFunc
}

type DoFunc func(ctx context.Context) error

func NewWork(do func(ctx context.Context) error) *Work {
	return &Work{
		do: do,
	}
}

func (w *Work) WithCond(cond CondFunc) *Work {
	cw := *w
	cw.cond = cond
	return &cw
}

func (w *Work) Do(ctx context.Context) error {
	cw := w
	for {
		if cw == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			ok, ctx, err := cw.ok(ctx)
			if err != nil {
				return err
			}
			if !ok {
				return nil
			}
			if cw.do != nil {
				if err := cw.do(ctx); err != nil {
					return err
				}
			}
			cw = cw.next
		}
	}
}

func (w *Work) ok(ctx context.Context) (bool, context.Context, error) {
	if w.cond == nil {
		return true, ctx, nil
	}
	return w.cond(ctx)
}
