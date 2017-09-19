package kuzzle_test

import (
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"fmt"
)

func ExampleKuzzle_Disconnect() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	err := k.Disconnect()
	if err != nil {
		fmt.Println(err.Error())
	}
}