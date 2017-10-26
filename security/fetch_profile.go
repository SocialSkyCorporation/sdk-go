package security

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) FetchProfile(id string, options types.QueryOptions) (*Profile, error) {
	res, err := s.rawFetch("getProfile", id, options)

	if err != nil {
		return nil, err
	}

	jsonProfile := &jsonProfile{}
	json.Unmarshal(res, jsonProfile)

	return s.jsonProfileToProfile(jsonProfile), nil
}