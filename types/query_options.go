package types

type QueryOptions interface {
	GetQueuable() bool
	SetQueuable(bool) *queryOptions
	GetFrom() int
	SetFrom(int) *queryOptions
	GetSize() int
	SetSize(int) *queryOptions
	GetScroll() string
	SetScroll(string) *queryOptions
	GetScrollId() string
	SetScrollId(string) *queryOptions
	GetVolatile() VolatileData
	SetVolatile(VolatileData) *queryOptions
	GetRefresh() string
	SetRefresh(string) *queryOptions
	GetIfExist() string
	SetIfExist(string) *queryOptions
}

type queryOptions struct {
	queuable bool
	from     int
	size     int
	scroll   string
	scrollId string
	volatile VolatileData
	refresh  string
	ifExist  string
}

func (qo queryOptions) GetQueuable() bool {
	return qo.queuable
}

func (qo *queryOptions) SetQueuable(queuable bool) *queryOptions {
	qo.queuable = queuable
	return qo
}

func (qo queryOptions) GetFrom() int {
	return qo.from
}

func (qo *queryOptions) SetFrom(from int) *queryOptions {
	qo.from = from
	return qo
}

func (qo queryOptions) GetSize() int {
	return qo.size
}

func (qo *queryOptions) SetSize(size int) *queryOptions {
	qo.size = size
	return qo
}

func (qo queryOptions) GetScroll() string {
	return qo.scroll
}

func (qo *queryOptions) SetScroll(scroll string) *queryOptions {
	qo.scroll = scroll
	return qo
}

func (qo queryOptions) GetScrollId() string {
	return qo.scrollId
}

func (qo *queryOptions) SetScrollId(scrollId string) *queryOptions {
	qo.scrollId = scrollId
	return qo
}

func (qo queryOptions) GetVolatile() VolatileData {
	return qo.volatile
}

func (qo *queryOptions) SetVolatile(volatile VolatileData) *queryOptions {
	qo.volatile = volatile
	return qo
}

func (qo queryOptions) GetRefresh() string {
	return qo.refresh
}

func (qo *queryOptions) SetRefresh(refresh string) *queryOptions {
	qo.refresh = refresh
	return qo
}

func (qo queryOptions) GetIfExist() string {
	return qo.ifExist
}

func (qo *queryOptions) SetIfExist(ifExist string) *queryOptions {
	qo.ifExist = ifExist
	return qo
}

func NewQueryOptions() *queryOptions {
	return &queryOptions{
		size:    10,
		ifExist: "error",
	}
}
