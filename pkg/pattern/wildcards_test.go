package pattern

import "testing"

func TestWithWildcards(t *testing.T) {
	cases := map[WithWildcards]map[string]bool{
		"a": {
			"a": true,
			"b": false,
			"*": false,
		},
		"abc": {
			"a":   false,
			"ab":  false,
			"abc": true,
		},
		"*": {
			"*":  true,
			"ab": true,
		},
		"a*b": {
			"ab":   true,
			"aab":  true,
			"aabb": true,
			"abab": true,
			"a":    false,
			"b":    false,
			"ba":   false,
		},
		"*a*b*": {
			"ab":   true,
			"xabx": true,
			"xax":  false,
			"xbx":  false,
		},
	}

	for pattern, expected := range cases {
		for target, shouldMatch := range expected {
			matches := pattern.Matches(target)
			if shouldMatch && !matches {
				t.Errorf("expected '%s' to match '%s'", pattern, target)
			}
			if !shouldMatch && matches {
				t.Errorf("expected '%s' not to match '%s'", pattern, target)
			}
		}
	}
}
