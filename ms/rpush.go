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

// Rpush appends the specified value at the end of a list,
// only if the key already exists and if it holds a list.
func (ms *Ms) Rpush(source string, values []string, options types.QueryOptions) (int, error) {
	if len(values) == 0 {
		return 0, types.NewError("Ms.Rpush: please provide at least one value", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Values []string `json:"values"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "rpush",
		Id:         source,
		Body:       &body{Values: values},
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
