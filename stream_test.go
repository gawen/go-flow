package flow

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseChannelInt(t *testing.T) {
	var someInt int

	s, err := ParseStream(make(chan int))
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, s.chElemType, reflect.TypeOf(someInt))
}

func TestParseArrayInt(t *testing.T) {
	var someInt int

	s, err := ParseStream([]int{1, 1, 2, 3, 5})
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, s.chElemType, reflect.TypeOf(someInt))
}

func TestParseStream(t *testing.T) {
	var someInt int

	s, err := ParseStream([]int{1, 1, 2, 3, 5})
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, s.chElemType, reflect.TypeOf(someInt))

	s2, err := ParseStream(s)
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, s.chElemType, reflect.TypeOf(someInt))
	assert.Equal(t, s, s2)
}

func TestParseError(t *testing.T) {
	_, err := ParseStream(42)
	assert.Error(t, err)

	_, err = ParseStream("foo")
	assert.Error(t, err)
}
