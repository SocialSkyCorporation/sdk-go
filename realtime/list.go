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

package realtime

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// List lists all subscriptions on all indexes and all collections.
func (r *Realtime) List(index string, collection string) (json.RawMessage, error) {
	if index == "" || collection == "" {
		return nil, types.NewError("Realtime.List: index and collection required", 400)
	}

	query := &types.KuzzleRequest{
		Controller: "realtime",
		Action:     "list",
		Index:      index,
		Collection: collection,
	}

	result := make(chan *types.KuzzleResponse)

	go r.k.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
