package hoge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	assert.Equal(t, "foo", Foo())
}
