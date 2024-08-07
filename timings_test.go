package harp

import (
	"bytes"
	"go/format"
	"strings"
	"testing"

	"github.com/daspoet/harp/pkg/ast"
	"github.com/daspoet/harp/pkg/instrument"
	"github.com/daspoet/harp/pkg/pattern"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimings_Instrument(t *testing.T) {
	toInstrument := `
    package main

    import (
        "os"
    )

    type Fish struct {}

    func (f Fish) Swim() {}
    func (f Fish) talk() {}

    type Duck struct {}

    func (d Duck) Swim() {}
    func (d Duck) talk() {}
    `

	timings := Timings{
		Targets: []instrument.Target{
			{Type: pattern.WithWildcards("Duck"), Visibility: ast.Public},
		},
	}

	var out bytes.Buffer

	err := instrument.InstrumentFile("", toInstrument, &out, timings)
	require.NoError(t, err)

	expected := `
    package main

    import (
        "fmt"
        "os"
        "time"
    )

    type Fish struct {}

    func (f Fish) Swim() {
    }

    func (f Fish) talk() {
    }

    type Duck struct {}

    func (d Duck) Swim() {
        _before := time.Now()
        defer func() {
            fmt.Println("Duck.Swim took", time.Since(_before))
        }()
    }
    func (d Duck) talk() {
    }
    `

	assert.Equal(t, withoutBlankLines(formatGo(t, expected)), withoutBlankLines(formatGo(t, out.String())))
}

func formatGo(t *testing.T, raw string) string {
	formatted, err := format.Source([]byte(raw))
	require.NoError(t, err)

	return string(formatted)
}

func withoutBlankLines(raw string) string {
	lines := strings.Split(raw, "\n")

	var out []string
	for _, line := range lines {
		if line != "" {
			out = append(out, line)
		}
	}
	return strings.Join(out, "\n")
}
