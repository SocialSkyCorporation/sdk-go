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

func TestZaddEmptyElements(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Zadd("foo", []*types.MSSortedSet{}, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.Zadd: please provide at least one element", fmt.Sprint(err))
}

func TestZaddError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	sortedSet := []*types.MSSortedSet{
		{Score: 10, Member: "bar"},
		{Score: 5, Member: "foo"},
	}

	_, err := memoryStorage.Zadd("foo", sortedSet, qo)

	assert.NotNil(t, err)
}

func TestZadd(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zadd", parsedQuery.Action)

			r, _ := json.Marshal(2)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	sortedSet := []*types.MSSortedSet{
		{Score: 10, Member: "bar"},
		{Score: 5, Member: "foo"},
	}

	res, _ := memoryStorage.Zadd("foo", sortedSet, qo)

	assert.Equal(t, 2, res)
}

func TestZaddWithOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zadd", parsedQuery.Action)
			assert.Equal(t, true, options.GetCh())
			assert.Equal(t, true, options.GetIncr())
			assert.Equal(t, true, options.GetNx())
			assert.Equal(t, true, options.GetXx())

			r, _ := json.Marshal(2)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	sortedSet := []*types.MSSortedSet{
		{Score: 10, Member: "bar"},
		{Score: 5, Member: "foo"},
	}

	qo.SetCh(true)
	qo.SetIncr(true)
	qo.SetNx(true)
	qo.SetXx(true)
	res, _ := memoryStorage.Zadd("foo", sortedSet, qo)

	assert.Equal(t, 2, res)
}

func ExampleMs_Zadd() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	sortedSet := []*types.MSSortedSet{
		{Score: 10, Member: "bar"},
		{Score: 5, Member: "foo"},
	}

	qo.SetCh(true)
	qo.SetIncr(true)
	qo.SetNx(true)
	qo.SetXx(true)
	res, err := memoryStorage.Zadd("foo", sortedSet, qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
