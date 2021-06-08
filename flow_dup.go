package flow

import (
	"reflect"
)

//go:generate go run flow_dupn_gen.go

func (f *Flow) DupN(
	chBuf int,
	numOuts int,
	in interface{},
) []*Stream {
	inStream := f.MustParseStream(in)
	chType := reflect.ChanOf(reflect.BothDir, inStream.chElemType)

	var insType []reflect.Type
	for i := -1; i < numOuts; i++ {
		insType = append(insType, chType)
	}

	kType := reflect.FuncOf(insType, nil, false)
	kVal := reflect.MakeFunc(kType, func(args []reflect.Value) []reflect.Value {
		inChVal := args[0]

		for {
			// get input value
			inVal, ok := inChVal.Recv()
			if !ok {
				break
			}

			// forward to all outputs
			for i := 1; i < len(args); i++ {
				outChVal := args[i]
				outChVal.Send(inVal)
			}
		}

		return nil
	})

	pipe := NewPipe(f, chBuf, []*Stream{inStream}, kType)
	pipe.assertOutputCount(numOuts)
	_ = pipe.Run(1, kVal)
	return pipe.outs
}
