package harp

import (
	"testing"

	"github.com/shoenig/test"
)

func TestHarp_Greet(t *testing.T) {
	test.Eq(t, "hello from harp", Greet())
}
