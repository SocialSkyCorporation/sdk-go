package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Pttl returns the remaining time to live of a key, in milliseconds.
func (ms Ms) Pttl(key string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.Pttl: key required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "pttl",
		Id:         key,
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
