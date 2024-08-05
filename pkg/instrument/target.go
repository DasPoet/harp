package instrument

import (
	goast "go/ast"

	"github.com/daspoet/harp/pkg/ast"
	"github.com/daspoet/harp/pkg/pattern"
)

type Target struct {
	Type   pattern.Pattern
	Method pattern.Pattern

	Visibility ast.Visibility
}

func (t Target) Accepts(node *goast.FuncDecl) bool {
	if !t.acceptsByVisibility(node) {
		return false
	}
	if !t.acceptsByTypename(node) {
		return false
	}
	if !t.acceptsByMethodName(node) {
		return false
	}
	return true
}

func (t Target) acceptsByVisibility(node *goast.FuncDecl) bool {
	return ast.NodeVisibility(node).Matches(t.Visibility)
}

func (t Target) acceptsByTypename(node *goast.FuncDecl) bool {
	typename := ast.ReceiverTypeName(node)
	if typename == nil {
		return false
	}

	if t.Type == nil {
		return true
	}
	return t.Type.Matches(*typename)
}

func (t Target) acceptsByMethodName(node *goast.FuncDecl) bool {
	if t.Method == nil {
		return true
	}
	return t.Method.Matches(node.Name.Name)
}
