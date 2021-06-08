# 🌊 `go-flow`

*🦋 work in progress*

Data processing as streams with a bit of Go's reflection.

```go
// Map cities to their weather using the wttr.in API.

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gawen/go-flow"
)

type WttrinInfo struct {
	city string

	CurrentCondition []struct {
		TemperatureC string `json:"temp_C"`
	} `json:"current_condition"`
}

func main() {
    // declare a flow
    f := flow.F(context.Background())

    // build a stream (`chan string`) with some cities.
    cities := f.Pipe(1, 0, func(cities chan string) {
        cities <- "Paris"
        cities <- "San Francisco"
        cities <- "New York"
        // ...
    })

    // get weather info for each city. Request wttr.in with 8 parallel lanes.
    wttrninInfos := cities.Pipe(8, 0, func(ctx context.Context, cities chan string, infos chan WttrinInfo) error {
        for city := range cities {
            url := fmt.Sprintf("https://wttr.in/%s?format=j1", city)
            req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
            if err != nil {
                return err
            }

            res, err := http.DefaultClient.Do(req)
            if err != nil {
                return err
            }
            defer res.Body.Close()

            buf, err := ioutil.ReadAll(res.Body)
            if err != nil {
                return err
            }

            var info WttrinInfo
            if err := json.Unmarshal(buf, &info); err != nil {
                return err
            }

            info.city = city
            infos <- info
        }

        return nil
    })

    // consume weather infos and print them.
    wttrninInfos.Consume(1, func(infos chan WttrinInfo) {
        for info := range infos {
            fmt.Printf("%+v\n", info)
        }
    })

    // wait for the flow to finish and panic if an error happened.
    if err := f.Wait(); err != nil {
        panic(err.Error())
    }
}
```

## Features

The core method is

```go
// Pipe `inputs` to `outputs` with `kernel`. Run parallely `laneCount` times.
func (*Flow) PipeN(
    laneCount int,          // instance count of running `kernel`. >=1.
    outChannelBuf int,      // outputs's channel buffer size. >=0.
    outsCount int,          // count of outputs. >=0.
    inputs ...interface{},  // list of inputs. can be channels, array/slice, or *flow.Stream
    kernel interface{},     // Kernel
) (outputs []*Stream)
```

A kernel is a method where the first arguments are the inputs given to `PipeN`,
and the rest the outputs. It can have an optional `context.Context` as first
argument, being the flow's `context.Context`.

```go
f.PipeN(
    8,                      // run 8 kernels simulatenously
    0,                      // outputs's channel won't have any buffer
    1,                      // only one output expected
    []int{1,2,3},           // first input: an []int. will be translated to a chan int.
    make(chan string),      // second input: an chan string.
    func(                   // define kernel
        inp1 chan int,      // first input
        inp2 chan string,   // second input
        out chan float,     // define one output, being a chan float
    ) {
        // consume inp1 and inp2 and output something in out.
    },
)
```

The inputs types are checked at runtime. If they do not match, `go-flow` will panic.

## Sugar

Sugar methods are provided to ease building flows.

- `Flow.Consume` calls `Flow.PipeN` with no expected output;
- `Flow.Pipe` does with one expected output;
- `Flow.Pipe2` does with 2 expected outputs;
- etc...

A `Stream` also provides the same methods, to chain calls and ease flow writing.
(see above example).

## Error handling

A kernel can return an optional error, which will kill the whole flow.

```go
f := flow.F(context.Background())

f.RangeInt(0, 10, 1).Pipe(1, 0, func(in, out chan int) error {
    for i := range in {
        // dummy error
        if i == 5 {
            return errors.New("i don't like 5s")
        }

        out <- i
    }

    return nil
}).Consume(1, func(in chan int) {
    for i := range in {
        println(i)
    }
})

if err := f.Wait(); err != nil {
    panic(err.Error())
}

// will print:
// 0
// 1
// 2
// 3
// 4
// panic: i don't like 5s
```
