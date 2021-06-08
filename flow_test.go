package flow

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleFlow(t *testing.T) {
	f := F(context.Background())
	assert.NotNil(t, f)

	m := make(map[int]bool)
	f.Consume(1, RangeInt(f, 0, 10, 1), func(is chan int) {
		for i := range is {
			m[i] = true
		}
	})

	assert.NoError(t, f.Wait())
	for i := 0; i < 10; i++ {
		if _, ok := m[i]; !ok {
			t.Errorf("missing %d", i)
		}
	}
}

func TestErrorFlow(t *testing.T) {
	someError := errors.New("some error")

	f := F(context.Background())
	assert.NotNil(t, f)

	m := make(map[int]bool)
	f.Consume(1, RangeInt(f, 0, 10, 1), func(is chan int) error {
		for i := range is {
			m[i] = true

			if i == 5 {
				return someError
			}
		}

		return nil
	})

	assert.Equal(t, someError, f.Wait())
	for i := 0; i <= 5; i++ {
		if _, ok := m[i]; !ok {
			t.Errorf("missing %d", i)
		}
	}

	for i := 6; i < 10; i++ {
		if _, ok := m[i]; ok {
			t.Errorf("having %d", i)
		}
	}
}

func TestCancelledFlow(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	f := F(ctx)
	assert.NotNil(t, f)

	m := make(map[int]bool)
	f.Consume(1, RangeInt(f, 0, 10, 1), func(is chan int) error {
		for i := range is {
			m[i] = true

			if i == 5 {
				cancel()
			}
		}

		return nil
	})

	err := f.Wait()
	assert.NotNil(t, ctx.Err())
	assert.Equal(t, ctx.Err(), err)
	for i := 0; i <= 5; i++ {
		if _, ok := m[i]; !ok {
			t.Errorf("missing %d", i)
		}
	}

	// context is done but cannot assert it stopped exactly at 5.
}
