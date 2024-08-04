package ast

import (
	"testing"

	"github.com/shoenig/test"
)

func TestVisibility_String(t *testing.T) {
	cases := map[Visibility]string{
		All:            "*",
		Private:        "private",
		Public:         "public",
		Visibility(42): "42",
	}

    for v, expected := range cases {
        test.Eq(t, expected, v.String())
    }
}
