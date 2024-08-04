package ast

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"

	"golang.org/x/tools/go/ast/astutil"
)

type ReplaceFunc func(node ast.Node) ast.Node

func Format(node ast.Node, fset *token.FileSet, out io.Writer) error {
	return printer.Fprint(out, token.NewFileSet(), node)
}

func ParseFile(filename string) (ast.Node, *token.FileSet, error) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}
	return node, fset, nil
}

func Replace(root ast.Node, fn ReplaceFunc) ast.Node {
	pre := func(c *astutil.Cursor) bool {
		var (
			original    = c.Node()
			replacement = fn(original)
		)
		if replacement != nil {
			c.Replace(replacement)
		}
		return true
	}
	return astutil.Apply(root, pre, nil)
}

func EnsureImport(node ast.Node, fset *token.FileSet, path string) {
	file, ok := node.(*ast.File)
	if !ok {
		panic("node is not a file")
	}
	astutil.AddImport(fset, file, path)
}
