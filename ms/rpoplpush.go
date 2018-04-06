// Copyright 2015-2017 Kuzzle
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

// RpoplPush removes the last element of the list at source and pushes it
// back at the start of the list at destination.
func (ms *Ms) Rpoplpush(source string, destination string, options types.QueryOptions) (*string, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Source      string `json:"source"`
		Destination string `json:"destination"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "rpoplpush",
		Body:       &body{Source: source, Destination: destination},
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return &returnedResult, nil
}
