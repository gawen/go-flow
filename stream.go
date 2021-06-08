package flow

import (
	"fmt"
	"reflect"
)

type Stream struct {
	f          *Flow
	chElemType reflect.Type
	chVal      reflect.Value
	plugged    bool
}

func ParseStream(x interface{}) (s *Stream, err error) {
	// is x already a Stream?
	if stream, ok := x.(*Stream); ok {
		s = stream
		return
	}

	xType := reflect.TypeOf(x)
	switch xType.Kind() {
	case reflect.Chan:
		s = &Stream{
			chElemType: xType.Elem(),
			chVal:      reflect.ValueOf(x),
		}

	case reflect.Array:
	case reflect.Slice:
		s = &Stream{}
		s.chElemType = xType.Elem()
		s.chVal = reflect.MakeChan(reflect.ChanOf(reflect.BothDir, s.chElemType), 1)

		go func() {
			defer func() {
				// only panic which can happen is if s.chVal is closed because an
				// error was raised or context was done.
				_ = recover()
			}()

			xVal := reflect.ValueOf(x)
			xLen := xVal.Len()
			for xIdx := 0; xIdx < xLen; xIdx++ {
				xIdxVal := xVal.Index(xIdx)
				s.chVal.Send(xIdxVal)
			}
		}()

	default:
		err = fmt.Errorf("expected an array, slice, chan or *flow.Stream, got %s", xType)
	}

	return
}

func MustParseStream(x interface{}) *Stream {
	stream, err := ParseStream(x)
	if err != nil {
		panic(err.Error())
	}

	return stream
}

func ParseStreams(xs ...interface{}) (r []*Stream, err error) {
	for xIdx, x := range xs {
		var stream *Stream
		stream, err = ParseStream(x)
		if err != nil {
			err = fmt.Errorf("arg #%d: %w", xIdx+1, err)
			break
		}

		r = append(r, stream)
	}

	return
}

func MustParseStreams(xs ...interface{}) []*Stream {
	streams, err := ParseStreams(xs...)
	if err != nil {
		panic(err.Error())
	}

	return streams
}

func (s *Stream) assertFlowAttached() {
	if s.f == nil {
		panic("no flow attached")
	}
}
