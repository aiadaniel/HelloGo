package engine

import (
	"fmt"
)

type SimpleEngine struct {
}

func (e *SimpleEngine) Run(isLocal bool, seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		result, err := Worker(r, isLocal)
		if err != nil {
			continue
		}
		requests = append(requests, result.Requests...)

		for _, item := range result.Items {
			fmt.Printf("Got item %v\n", item)
		}
	}
}
