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
	f, err := os.Create("flow_dupn.go")
	die(err)
	defer f.Close()

	fmt.Fprintf(f, `// Code generated by go generate; DO NOT EDIT.
// This file was generated by flow_dupn_gen.go at %s
package flow

`, time.Now().UTC())

	for i := 2; i <= PipeXCount; i++ {
		fmt.Fprintf(f, "func (f *Flow) Dup%d(chBuf int, in interface{}) (", i)

		{
			var outs []string
			for j := 0; j < i; j++ {
				outs = append(outs, "*Stream")
			}
			fmt.Fprintf(f, strings.Join(outs, ", "))
		}

		fmt.Fprintf(f, `) {
	outs := f.DupN(chBuf, %d, in)
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
