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
	switch strings.ToLower(raw) {
	case "*":
		return All, nil
	case "private":
		return Private, nil
	case "public":
		return Public, nil
	}
	return 0, errors.New("invalid visibility: '" + raw + "', expected one of: 'all', 'private', 'public'")
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
