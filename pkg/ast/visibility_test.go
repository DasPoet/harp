package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVisibility_String(t *testing.T) {
	cases := map[Visibility]string{
		All:            "*",
		Private:        "private",
		Public:         "public",
		Visibility(42): "42",
	}

    for v, expected := range cases {
        assert.Equal(t, expected, v.String())
    }
}
