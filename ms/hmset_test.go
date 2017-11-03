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

func TestHmsetEmptyEntries(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.MemoryStorage.Hmset("", []*types.MsHashField{}, nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.Hmset: at least one entry field to set is required", fmt.Sprint(err))
}

func TestHmsetError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.Hmset("foo", []*types.MsHashField{}, nil)

	assert.NotNil(t, err)
}

func TestHmset(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "hmset", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)
			assert.Equal(t, "foo", parsedQuery.Body.(map[string]interface{})["entries"].([]interface{})[0].(map[string]interface{})["field"].(string))
			assert.Equal(t, "bar", parsedQuery.Body.(map[string]interface{})["entries"].([]interface{})[0].(map[string]interface{})["value"].(string))

			r, _ := json.Marshal("result")
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.MemoryStorage.Hmset("foo", []*types.MsHashField{{Field: "foo", Value: "bar"}}, nil)

	assert.Equal(t, "result", res)
}

func ExampleMs_Hmset() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := k.MemoryStorage.Hmset("foo", []*types.MsHashField{{Field: "foo", Value: "bar"}}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
