package pattern

// WithWildcards is a string that can contain wildcards '*'.
type WithWildcards string

// Matches reports whether w matches target. Wildcards
// in target are treated as regular characters.
func (w WithWildcards) Matches(target string) bool {
	var (
		targetIndex  int
		patternIndex int

		lastStarIndex  = -1
		backtrackIndex = -1
	)

	for targetIndex < len(target) {
		if patternIndex < len(w) && w[patternIndex] == target[targetIndex] {
			// match
			targetIndex++
			patternIndex++
		} else if patternIndex < len(w) && w[patternIndex] == '*' {
			// wildcard
			lastStarIndex = patternIndex
			backtrackIndex = targetIndex
			patternIndex++
		} else if lastStarIndex != -1 {
			// backtrack - there's a previous wildcard
			patternIndex = lastStarIndex + 1
			backtrackIndex++
			targetIndex = backtrackIndex
		} else {
			// mo match and no previous wildcard
			return false
		}
	}
	return w[patternIndex:].containsOnly('*')
}

func (w WithWildcards) containsOnly(r rune) bool {
	for _, c := range w {
		if c != r {
			return false
		}
	}
	return true
}

var _ Pattern = (*WithWildcards)(nil)
