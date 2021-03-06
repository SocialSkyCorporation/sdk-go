// Copyright 2015-2018 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Zremrangebylex removes members from a sorted set where all elements have the same score, using lexicographical ordering. The min and max interval are inclusive, see the Redis documentation to change this behavior.
func (ms *Ms) Zremrangebylex(key string, min string, max string, options types.QueryOptions) (int, error) {
	if min == "" || max == "" {
		return 0, types.NewError("Ms.Zremrangebylex: an empty string is not a valid string range item", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Min string `json:"min"`
		Max string `json:"max"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zremrangebylex",
		Id:         key,
		Body:       &body{Min: min, Max: max},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return 0, res.Error
	}

	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
