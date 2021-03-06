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
	"strconv"
)

// Zrevrangebyscore is identical to zrangebyscore except that the sorted set is traversed in descending order.
func (ms *Ms) Zrevrangebyscore(key string, min float64, max float64, options types.QueryOptions) ([]*types.MSSortedSet, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zrevrangebyscore",
		Id:         key,
		Min:        strconv.FormatFloat(min, 'f', 6, 64),
		Max:        strconv.FormatFloat(max, 'f', 6, 64),
	}

	assignZrangeOptions(query, options)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return mapZrangeResults(returnedResult)
}
