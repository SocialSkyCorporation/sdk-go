package types

type roomOptions struct {
	scope           string
	state           string
	user            string
	subscribeToSelf bool
}

type RoomOptions interface {
	GetScope() string
	SetScope(string)
	GetState() string
	SetState(string)
	GetUser() string
	SetUser(string)
	GetSubscribeToSelf() bool
	SetSubscribeToSelf(bool)
}

func (ro roomOptions) GetScope() string {
	return ro.scope
}

func (ro *roomOptions) SetScope(scope string) *roomOptions {
	ro.scope = scope
	return ro
}

func (ro roomOptions) GetState() string {
	return ro.state
}

func (ro *roomOptions) SetState(state string) *roomOptions {
	ro.state = state
	return so
}

func (ro roomOptions) GetUser() string {
	return ro.user
}

func (ro *roomOptions) SetUser(user string) *roomOptions {
	ro.user = user
	return ro
}

func (ro roomOptions) GetSubscribeToSelf() bool {
	return ro.subscribeToSelf
}

func (ro *roomOptions) SetSubscribeToSelf(subscribeToSelf bool) *roomOptions {
	ro.subscribeToSelf = subscribeToSelf
	return ro
}

func NewRoomOptions() *roomOptions {
	return &roomOptions{
		scope:           SCOPE_ALL,
		state:           STATE_DONE,
		user:            USER_NONE,
		subscribeToSelf: true,
	}
}
