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

func TestZremRangeByLexEmptyMin(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)

	_, err := memoryStorage.ZremRangeByLex("foo", "", "(f", nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.ZremRangeByLex: an empty string is not a valid string range item", fmt.Sprint(err))
}

func TestZremRangeByLexEmptyMax(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)

	_, err := memoryStorage.ZremRangeByLex("foo", "[b", "", nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.ZremRangeByLex: an empty string is not a valid string range item", fmt.Sprint(err))
}

func TestZremRangeByLexError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.ZremRangeByLex("foo", "[b", "(f", qo)

	assert.NotNil(t, err)
}

func TestZremRangeByLex(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zremrangebylex", parsedQuery.Action)

			r, _ := json.Marshal(42)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	res, _ := memoryStorage.ZremRangeByLex("foo", "[b", "(f", qo)

	assert.Equal(t, 42, res)
}

func ExampleMs_ZremRangeByLex() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	res, err := memoryStorage.ZremRangeByLex("foo", "[b", "(f", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
