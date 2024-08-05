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

func TestVisibility_Matches(t *testing.T) {
    cases := map[Visibility]map[Visibility]bool{
        All: {
            All:     true,
            Private: true,
            Public:  true,
        },
        Private: {
            All:     true,
            Private: true,
            Public:  false,
        },
        Public: {
            All:     true,
            Private: false,
            Public:  true,
        },
    }

    for v, expected := range cases {
        for other, shouldMatch := range expected {
            assert.Equal(t, shouldMatch, v.Matches(other), "expected '%s' to match '%s'", v, other)
        }
    }
}
