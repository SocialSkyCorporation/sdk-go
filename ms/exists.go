package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Exists checks if the specified keys exist
func (ms *Ms) Exists(keys []string, options types.QueryOptions) (int, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Keys []string `json:"keys"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "exists",
		Body:       &body{Keys: keys},
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
