package main

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
	defer println("exited")

	f := flow.F(context.Background())

	f.Pipe(100, 0, func(cities chan string) {
		cities <- "Paris"
		cities <- "San Francisco"
		cities <- "New York"
	}).Pipe(16, 0, func(ctx context.Context, cities chan string, infos chan WttrinInfo) error {
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
	}).Consume(1, func(infos chan WttrinInfo) {
		for info := range infos {
			fmt.Printf("%+v\n", info)
		}
	})

	if err := f.Wait(); err != nil {
		panic(err.Error())
	}
}
