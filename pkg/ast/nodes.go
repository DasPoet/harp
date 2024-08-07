package ast

import (
	goast "go/ast"
)

func NodeVisibility(node *goast.FuncDecl) Visibility {
	if !node.Name.IsExported() {
		return Private
	}
	return Public
}

func ReceiverTypeName(node *goast.FuncDecl) *string {
	recv := node.Recv.List
	if len(recv) != 1 {
		return nil
	}

	ident, ok := recv[0].Type.(*goast.Ident)
	if !ok {
		return nil
	}
	return &ident.Name
}
