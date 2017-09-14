package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

// ZincrBy increments the score of a member in a sorted set by the provided value.
func (ms Ms) ZincrBy(key string, member string, increment float64, options types.QueryOptions) (float64, error) {
	if key == "" {
		return 0, errors.New("Ms.ZincrBy: key required")
	}
	if member == "" {
		return 0, errors.New("Ms.ZincrBy: member required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Member    string  `json:"member"`
		Increment float64 `json:"value"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "zincrby",
		Id:         key,
		Body:       &body{Member: member, Increment: increment},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}

	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)
	parsedResult, _ := strconv.ParseFloat(returnedResult, 64)

	return parsedResult, nil
}
