package collection

import (
	"container/list"
	"encoding/json"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/satori/go.uuid"
)

type Room struct {
	RequestId       string           `json:"RequestId"`
	RoomId          string           `json:"roomId"`
	Channel         string           `json:"channel"`
	result          json.RawMessage  `json:"-"`
	scope           string           `json:"-"`
	state           string           `json:"-"`
	user            string           `json:"-"`
	subscribeToSelf bool             `json:"-"`

	collection                  *Collection                      `json:"-"`
	RealtimeNotificationChannel chan<- *types.KuzzleNotification `json:"-"`
	subscribeResponseChan       chan<- *types.SubscribeResponse

	pendingSubscriptions map[string]chan<- *types.KuzzleNotification `json:"-"`
	subscribing          bool                                        `json:"-"`
	queue                *list.List                                  `json:"-"`
	id                   string                                      `json:"-"`
	Volatile             types.VolatileData                          `json:"-"`
	filters              interface{}                                 `json:"-"`
}

// NewRoom instanciates a new Room; this type is the result of a subscription request,
// allowing to manipulate the subscription itself.
// In Kuzzle, you don't exactly subscribe to a room or a topic but, instead, you subscribe to documents.
// What it means is that, to subscribe, you provide to Kuzzle a set of matching filters.
// Once you have subscribed, if a pub/sub message is published matching your filters, or if a matching stored
// document change (because it is created, updated or deleted), then you'll receive a notification about it.
func NewRoom(c *Collection, opts types.RoomOptions) *Room {
	if opts == nil {
		opts = types.NewRoomOptions()
	}
	r := &Room{
		scope:                opts.GetScope(),
		state:                opts.GetState(),
		user:                 opts.GetUser(),
		id:                   uuid.NewV4().String(),
		collection:           c,
		pendingSubscriptions: make(map[string]chan<- *types.KuzzleNotification),
		subscribeToSelf:      opts.GetSubscribeToSelf(),
		Volatile:             opts.GetVolatile(),
		queue:								&list.List{},
	}
	r.queue.Init()

	return r
}

// GetRealtimeChannel return the room's ReatimeNotificationChannel
func (room Room) GetRealtimeChannel() chan<- *types.KuzzleNotification {
	return room.RealtimeNotificationChannel
}

// isReady returns true if the room is ready
func (room Room) isReady() bool {
	return *room.collection.Kuzzle.State == state.Connected && !room.subscribing
}

// GetRoomId returns the room's id
func (room Room) GetRoomId() string {
	return room.RoomId
}

// GetFilters returns the room's filters
func (room Room) GetFilters() interface{} {
	return room.filters
}

// GetResponseChannel returns the room's response channel
func (room Room) GetResponseChannel() chan<- *types.SubscribeResponse {
	return room.subscribeResponseChan
}
