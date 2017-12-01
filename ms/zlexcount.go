package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// ZlexCount counts elements in a sorted set where all members have equal score, using lexicographical ordering. The min and max values are inclusive by default. To change this behavior, please check the syntax detailed in the Redis documentation.
func (ms *Ms) Zlexcount(key string, min string, max string, options types.QueryOptions) (int, error) {
	if min == "" || max == "" {
		return 0, types.NewError("Ms.Zlexcount: an empty string is not a valid string range item", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zlexcount",
		Id:         key,
		Min:        min,
		Max:        max,
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
