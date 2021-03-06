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
	"fmt"

	"github.com/kuzzleio/sdk-go/types"
)

// ValidateSpecifications validates the provided specifications.
func (dc *Collection) ValidateSpecifications(index string, collection string, specifications json.RawMessage, options types.QueryOptions) (*types.ValidationResponse, error) {
	if index == "" {
		return nil, types.NewError("Collection.ValidateSpecifications: index required", 400)
	}

	if collection == "" {
		return nil, types.NewError("Collection.ValidateSpecifications: collection required", 400)
	}

	if specifications == nil {
		return nil, types.NewError("Collection.ValidateSpecifications: specifications required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	body := make(map[string]map[string]json.RawMessage)
	body[index] = make(map[string]json.RawMessage)
	body[index][collection] = specifications

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return nil, types.NewError(fmt.Sprintf("Unable to construct body: %s\n", err.Error()), 500)
	}

	query := &types.KuzzleRequest{
		Controller: "collection",
		Action:     "validateSpecifications",
		Body:       json.RawMessage(jsonBody),
	}

	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	vr, err := types.NewValidationResponse(res.Result)

	if err != nil {
		return nil, err
	}

	return vr, nil
}
