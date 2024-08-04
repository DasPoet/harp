package ast

import "strconv"

type Visibility int

const (
	All Visibility = iota
	Private
	Public
)

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
