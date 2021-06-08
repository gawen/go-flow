package flow

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var f = F(context.Background())

func emptyInt(f *Flow) *Stream {
	return f.Pipe(1, 0, func(o chan int) {})
}

func emptyString(f *Flow) *Stream {
	return f.Pipe(1, 0, func(o chan string) {})
}

func TestBadInput(t *testing.T) {
	assert.Panics(t, func() {
		f.Consume(0, func() {})
	}, "no lane")

	assert.Panics(t, func() {
		f.Consume(0, func(o chan int) {})
	}, "expected 0 outputs, got 1")

	assert.Panics(t, func() {
		f.Consume(0, emptyInt(f), func(i, o chan int) {})
	}, "expected 0 outputs, got 1")

	assert.Panics(t, func() {
		f.Consume(0, emptyInt(f), emptyInt(f), func(i chan int) {})
	}, "expected 2 inputs, got 1")

	assert.Panics(t, func() {
		f.Consume(0, emptyString(f), func(i chan int) {})
	}, "expected a chan int at argument 1, got chan string")
}
