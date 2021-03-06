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

package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Create creates a new empty data collection
func (dc *Collection) Create(index string, collection string, body json.RawMessage, options types.QueryOptions) error {
	if index == "" {
		return types.NewError("Collection.Create: index required", 400)
	}

	if collection == "" {
		return types.NewError("Collection.Create: collection required", 400)
	}

	if body != nil && !json.Valid(body) {
		return types.NewError("Collection.Create: body is not a valid JSON object", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "collection",
		Action:     "create",
		Body:       body,
	}

	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() == "" {
		return nil
	}

	return res.Error
}
