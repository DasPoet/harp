package harp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHarp_Greet(t *testing.T) {
    assert.Equal(t, "hello from harp", Greet())
}
