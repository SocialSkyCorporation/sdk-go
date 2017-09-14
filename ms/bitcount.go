package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Bitcount counts the number of set bits (population counting)
func (ms Ms) Bitcount(key string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Bitcount: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "bitcount",
		Id:         key,
	}

	if options != nil {
		if options.GetStart() != 0 {
			var start = options.GetStart()
			query.Start = &start
		}

		if options.GetEnd() != 0 {
			query.End = options.GetEnd()
		}
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
