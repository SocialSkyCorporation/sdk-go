package ms_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeoposError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.Geopos("foo", []string{}, nil)

	assert.NotNil(t, err)
}

func TestGeopos(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "geopos", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)
			assert.Equal(t, []string{"some", "members"}, parsedQuery.Members)

			r, _ := json.Marshal([][]string{{"43.6075274", "3.9128795"}, {"25.176", "14.466577"}})
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.MemoryStorage.Geopos("foo", []string{"some", "members"}, nil)

	assert.Equal(t, []*types.GeoPoint{{float64(43.6075274), float64(3.9128795), "some"}, {float64(25.176), float64(14.466577), "members"}}, res)
}

func TestGeoposLonConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "geopos", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)
			assert.Equal(t, []string{"members"}, parsedQuery.Members)

			r, _ := json.Marshal([][]string{{"43.6075abc", "3.9128795"}})
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.Geopos("foo", []string{"members"}, nil)

	assert.NotNil(t, err)
}

func TestGeoposLatConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "geopos", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)
			assert.Equal(t, []string{"members"}, parsedQuery.Members)

			r, _ := json.Marshal([][]string{{"43.6075274", "3.9128abc"}})
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.Geopos("foo", []string{"members"}, nil)

	assert.NotNil(t, err)
}

func ExampleMs_Geopos() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := k.MemoryStorage.Geopos("foo", []string{"members"}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Name, res[0].Lon, res[0].Lat)
}
