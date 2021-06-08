package flow

import (
	"reflect"
)

//go:generate go run flow_pipen_gen.go

func parsePipeArgs(args []interface{}) (inps []interface{}, k interface{}) {
	if len(args) < 1 {
		panic("expected kernel, got none")
	}

	inps = args[:len(args)-1]
	k = args[len(args)-1]
	return
}

func (f *Flow) PipeReflect(
	lane int,
	chBuf int,
	inps []interface{},
	kType reflect.Type,
	kVal reflect.Value,
) []*Stream {
	pipe := NewPipe(f, chBuf, f.MustParseStreams(inps...), kType)
	_ = pipe.Run(lane, kVal)
	return pipe.outs
}

func (f *Flow) PipeN(lane, chBuf int, numOuts int, args ...interface{}) []*Stream {
	inps, k := parsePipeArgs(args)
	pipe := NewPipe(f, chBuf, f.MustParseStreams(inps...), reflect.TypeOf(k))
	pipe.assertOutputCount(numOuts)
	_ = pipe.Run(lane, reflect.ValueOf(k))
	return pipe.outs
}
