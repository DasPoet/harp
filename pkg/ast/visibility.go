package ast

import (
	"errors"
	"strconv"
	"strings"
)

type Visibility int

const (
	All Visibility = iota
	Private
	Public
)

func ParseVisibility(raw string) (Visibility, error) {
	visibilityByRaw := map[string]Visibility{
		"*":       All,
		"private": Private,
		"public":  Public,
	}

    visibility, ok := visibilityByRaw[strings.ToLower(raw)]
    if !ok {
        return 0, errors.New("invalid visibility: '" + raw + "', expected one of: " + strings.Join(keys(visibilityByRaw), ", "))
    }
    return visibility, nil
}

func (v Visibility) Matches(other Visibility) bool {
	if v == All || other == All {
		return true
	}
	return v == other
}

func (v Visibility) String() string {
	switch v {
	case All:
		return "*"
	case Private:
		return "private"
	case Public:
		return "public"
	}
	return strconv.Itoa(int(v))
}

func keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}
