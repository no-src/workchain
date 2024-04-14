package workchain

import "context"

type Work struct {
	prev *Work
	next *Work

	do func(ctx context.Context) error
}

func NewWork(do func(ctx context.Context) error) *Work {
	return &Work{
		do: do,
	}
}

func (w *Work) Do(ctx context.Context) (err error) {
	cw := w
	for {
		if cw == nil {
			return err
		}
		if cw.do != nil {
			if err = cw.do(ctx); err != nil {
				return err
			}
		}
		cw = cw.next
	}
}
