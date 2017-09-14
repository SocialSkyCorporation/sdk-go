package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Zrem removes members from a sorted set.
func (ms Ms) Zrem(key string, members []string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Zrem: key required")
	}
	if len(members) == 0 {
		return 0, errors.New("Ms.Zrem: please provide at least one member")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Members []string `json:"members"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "zrem",
		Id:         key,
		Body:       &body{Members: members},
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
