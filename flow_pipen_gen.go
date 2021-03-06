// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	PipeXCount = 16
)

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	f, err := os.Create("flow_pipen.go")
	die(err)
	defer f.Close()

	fmt.Fprintf(f, `// Code generated by go generate; DO NOT EDIT.
// This file was generated by flow_pipen_gen.go at %s
package flow

`, time.Now().UTC())

	fmt.Fprintf(f, `func (f *Flow) Consume(lane int, args ...interface{}) {
	_ = f.PipeN(lane, 0, 0, args...)
}

`)

	for i := 1; i <= PipeXCount; i++ {
		fmt.Fprintf(f, "func (f *Flow) Pipe")
		if i > 1 {
			fmt.Fprintf(f, "%d", i)
		}
		fmt.Fprintf(f, "(lane, chBuf int, args ...interface{}) ")

		if i > 1 {
			fmt.Fprintf(f, "(")
		}

		{
			var outs []string
			for j := 0; j < i; j++ {
				outs = append(outs, "*Stream")
			}
			fmt.Fprintf(f, strings.Join(outs, ", "))
		}

		if i > 1 {
			fmt.Fprintf(f, ")")
		}

		fmt.Fprintf(f, ` {
	outs := f.PipeN(lane, chBuf, %d, args...)
	return `, i)
		{
			var outs []string
			for j := 0; j < i; j++ {
				outs = append(outs, fmt.Sprintf("outs[%d]", j))
			}
			fmt.Fprintf(f, strings.Join(outs, ", "))
		}
		fmt.Fprintf(f, `
}

`)

	}

}
