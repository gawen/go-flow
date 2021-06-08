package flow

import (
	"reflect"
	"sync"
)

type Pipe struct {
	f            *Flow
	argsVal      []reflect.Value
	outsVal      []reflect.Value
	outs         []*Stream
	returnsError bool
}

func NewPipe(
	f *Flow,
	chBuf int,
	inps []*Stream,
	kType reflect.Type,
) (p *Pipe) {

	// make sure it is indeed a kernel
	if kType.Kind() != reflect.Func {
		panicf("expected a function, got %s", kType)
	}

	p = &Pipe{}
	p.f = f

	// parse arguments
	inpIdx := 0
	for argIdx := 0; argIdx < kType.NumIn(); argIdx++ {
		argType := kType.In(argIdx)

		// argument has to be a channel, except the first one which can be a
		// `context.Context`.
		if argType.Kind() != reflect.Chan {
			// is a context?
			if argIdx == 0 {
				p.argsVal = append(p.argsVal, p.f.ctxVal)
				continue
			}

			panicf("expected a channel at argument %d, got %s", argIdx+1, argType)
		}

		if inpIdx < len(inps) { // are they inputs?
			if inps[inpIdx].plugged {
				panicf("arg #%d: stream has been already plugged. you may want to use Dup*().", argIdx+1)
			}
			inps[inpIdx].plugged = true

			// make sure the expected channel type is given
			if argType.Elem() != inps[inpIdx].chElemType {
				panicf("expected a %s at argument %d, got chan %s", argType, argIdx+1, inps[inpIdx].chElemType)
			}

			p.argsVal = append(p.argsVal, inps[inpIdx].chVal)
			inpIdx++
		} else { // they are outputs
			// build channel
			chVal := reflect.MakeChan(argType, chBuf)
			p.outsVal = append(p.outsVal, chVal)
			p.argsVal = append(p.argsVal, chVal)

			p.outs = append(p.outs, &Stream{
				f:          p.f,
				chElemType: argType.Elem(),
				chVal:      chVal,
			})
		}
	}

	if inpIdx != len(inps) {
		panicf("expected %d inputs, got %d", len(inps), inpIdx)
	}

	// does the kernel returns an error?
	if kType.NumOut() > 1 {
		panicf("expected kernel with one or no return value, got one with %d return values", kType.NumOut())
	}
	p.returnsError = kType.NumOut() == 1
	// XXX TODO make sure it returns an error

	return
}

func (p *Pipe) Run(
	lane int,
	kVal reflect.Value,
) chan struct{} {
	if lane < 1 {
		panic("lane must be at least 1")
	}

	var wg sync.WaitGroup
	wg.Add(lane)
	p.f.wg.Add(lane)

	// run core method
	for laneIdx := 0; laneIdx < lane; laneIdx++ {
		go func() {
			defer p.f.wg.Done()
			defer wg.Done()

			defer func() {
				// if context is done, the output channels are closed. If they
				// were being sent, a panic is raised that is caught. The
				// drawback is panics happening after a done context will be
				// panic-ed back, which will mess up a bit the panic stacktrace.

				// XXX if context is Done, just silent a panic?

				if p.f.ctx.Err() == nil {
					return
				}

				// context is done. recover panic
				if errI := recover(); errI != nil {
					// check panic is about a "send on closed channel"
					if err, ok := errI.(error); ok {
						if err.Error() == "send on closed channel" {
							// ignore panic.
							return
						}
					}

					// re-panic the original error
					panic(errI)
				}

			}()

			retsVal := kVal.Call(p.argsVal)

			if p.returnsError {
				errI := retsVal[0].Interface()
				if errI != nil {
					p.f.raise(errI.(error))
				}
			}
		}()
	}

	// closed when wg is done
	wgDoneCh := make(chan struct{})
	go func() {
		defer close(wgDoneCh)
		wg.Wait()
	}()

	// close outputs either when all is finished or context is done
	go func() {
		select {
		case <-wgDoneCh:
			break
		case <-p.f.ctx.Done():
			break
		}

		for _, outVal := range p.outsVal {
			outVal.Close()
		}
	}()

	return wgDoneCh
}

func (p *Pipe) assertOutputCount(numOuts int) {
	if numOuts >= 0 && len(p.outs) != numOuts {
		panicf("expected %v outputs, got %v", numOuts, len(p.outs))
	}
}
