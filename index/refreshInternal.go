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

package index

import "github.com/kuzzleio/sdk-go/types"

// RefreshInternal action forces a refresh, on the internal index,
// making the documents available to search immediately.
func (i *Index) RefreshInternal(options types.QueryOptions) error {

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "index",
		Action:     "refreshInternal",
	}
	go i.kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return res.Error
	}
	return nil
}
