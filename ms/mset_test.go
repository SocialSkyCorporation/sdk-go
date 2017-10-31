package ms_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	MemoryStorage "github.com/kuzzleio/sdk-go/ms"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMsetEmptyEntries(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	_, err := memoryStorage.Mset([]*types.MSKeyValue{}, nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.Mset: please provide at least one key/value entry", fmt.Sprint(err))
}

func TestMsetError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	entries := []*types.MSKeyValue{
		{Key: "foo", Value: "bar"},
		{Key: "bar", Value: "foo"},
	}

	_, err := memoryStorage.Mset(entries, qo)

	assert.NotNil(t, err)
}

func TestMset(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "mset", parsedQuery.Action)

			r, _ := json.Marshal("OK")
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	entries := []*types.MSKeyValue{
		{Key: "foo", Value: "bar"},
		{Key: "bar", Value: "foo"},
	}

	res, _ := memoryStorage.Mset(entries, qo)

	assert.Equal(t, "OK", res)
}

func ExampleMs_Mset() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	entries := []*types.MSKeyValue{
		{Key: "foo", Value: "bar"},
		{Key: "bar", Value: "foo"},
	}

	res, err := memoryStorage.Mset(entries, qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
