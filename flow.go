package flow

import (
	"context"
	"reflect"
	"sync"
)

type Flow struct {
	mu sync.Mutex

	// Context to which the flow is binded.
	// If context is done, flow will be cancelled and returns `context.Err()`.
	ctx context.Context

	// Method to cancel the flow's context.
	cancel func()

	// Reflect value of `ctx`. `= reflect.ValueOf(ctx)`
	ctxVal reflect.Value

	// Working group, done wait all kernels finished
	wg sync.WaitGroup

	// Error raised by one pipe in the flow.
	// protected by `mu`.
	err error
}

func NewFlow(ctx context.Context) *Flow {
	subCtx, cancel := context.WithCancel(ctx)

	return &Flow{
		ctx:    subCtx,
		cancel: cancel,
		ctxVal: reflect.ValueOf(ctx),
	}
}

var F = NewFlow

func (f *Flow) raise(err error) {
	// store the error
	f.mu.Lock()
	f.err = err
	f.mu.Unlock()

	// cancel the flow's context
	f.cancel()
}

// Returns the flow's error, or `nil` if none.
func (f *Flow) getError() error {
	// if an error was set, returns it
	f.mu.Lock()
	err := f.err
	f.mu.Unlock()
	if err != nil {
		return err
	}

	// else, returns the flow's context error; potentially `nil`.
	return f.ctx.Err()
}

func (f *Flow) Wait() error {
	f.wg.Wait()
	return f.getError()
}

func (f *Flow) MustParseStream(x interface{}) (stream *Stream) {
	stream = MustParseStream(x)
	stream.f = f
	return
}

func (f *Flow) MustParseStreams(xs ...interface{}) (streams []*Stream) {
	streams = MustParseStreams(xs...)
	for i := range streams {
		streams[i].f = f
	}
	return
}
