package collection_test

import (
	"testing"
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/stretchr/testify/assert"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

func TestSetHeaders(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(internal.MockedConnection{}, nil)

	m := make(map[string]interface{})
	m["foo"] = "bar"

	collection.NewCollection(k, "collection", "index").SetHeaders(m, false)

	assert.Equal(t, "bar", k.GetHeader("foo"))
}

func TestSetHeadersReplace(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(internal.MockedConnection{}, nil)

	m := make(map[string]interface{})
	m["foo"] = "bar"

	collection.NewCollection(k, "collection", "index").SetHeaders(m, false)

	assert.Equal(t, "bar", k.GetHeader("foo"))

	delete(m, "foo")
	m["oof"] = "bar"

	collection.NewCollection(k, "collection", "index").SetHeaders(m, true)

	assert.Nil(t, k.GetHeader("foo"))
	assert.Equal(t, "bar", k.GetHeader("oof"))
}
