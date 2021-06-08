package flow

import (
	"context"
)

func RangeInt(f *Flow, start, end, step int) *Stream {
	return f.Pipe(1, 0, func(ctx context.Context, o chan int) {
		for i := start; i < end; i = i + step {
			select {
			case o <- i:
			case <-ctx.Done():
				return
			}
		}
	})
}
