package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

type ZScanResponse struct {
	Cursor int
	Values []string
}

// Zscan is identical to scan, except that zscan iterates the members held by a sorted set.
func (ms Ms) Zscan(key string, cursor int, options types.QueryOptions) (*types.MSScanResponse, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zscan",
		Id:         key,
		Cursor:     cursor,
	}

	if options != nil {
		if options.GetCount() != 0 {
			query.Count = options.GetCount()
		}

		if options.GetMatch() != "" {
			query.Match = options.GetMatch()
		}
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	var scanResponse []interface{}
	json.Unmarshal(res.Result, &scanResponse)

	return formatZscanResponse(scanResponse), nil
}

func formatZscanResponse(response []interface{}) *types.MSScanResponse {
	formatedResponse := &types.MSScanResponse{}

	for _, element := range response {
		switch vf := element.(type) {
		case string:
			formatedResponse.Cursor, _ = strconv.Atoi(vf)
		case []interface{}:
			values := make([]string, 0, len(vf))

			for _, v := range vf {
				switch vv := v.(type) {
				case string:
					values = append(values, vv)
				}
			}

			formatedResponse.Values = values
		}
	}

	return formatedResponse
}
