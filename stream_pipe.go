package flow

import (
	"reflect"
)

//go:generate go run stream_pipen_gen.go

func (s *Stream) PipeReflect(
	lane int,
	chBuf int,
	inps []interface{},
	kType reflect.Type,
	kVal reflect.Value,
) []*Stream {
	s.assertFlowAttached()

	pipe := NewPipe(s.f, chBuf, s.f.MustParseStreams(inps...), kType)
	_ = pipe.Run(lane, kVal)
	return pipe.outs
}

func (s *Stream) PipeN(lane, chBuf int, numOuts int, k interface{}) []*Stream {
	inps := []interface{}{s}
	pipe := NewPipe(s.f, chBuf, s.f.MustParseStreams(inps...), reflect.TypeOf(k))
	pipe.assertOutputCount(numOuts)
	_ = pipe.Run(lane, reflect.ValueOf(k))
	return pipe.outs
}
