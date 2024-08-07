package instrument

import (
	goast "go/ast"
	"io"

	"github.com/daspoet/harp/pkg/ast"
)

type Context struct {
	importsToAdd []string
}

func (c *Context) AddImports(packages ...string) {
	c.importsToAdd = append(c.importsToAdd, packages...)
}

type Instrumentor interface {
	Instrument(ctx *Context, node goast.Node) goast.Node
}

func InstrumentFile(filename string, src any, out io.Writer, inst Instrumentor) error {
	node, fset, err := ast.ParseFile(filename, src)
	if err != nil {
		return err
	}

	ctx := Context{importsToAdd: []string{}}
	replaced := ast.Replace(
		node,
		func(node goast.Node) goast.Node {
			return inst.Instrument(&ctx, node)
		},
	)

	for _, imp := range ctx.importsToAdd {
		ast.EnsureImport(replaced, fset, imp)
	}
	return ast.Format(replaced, fset, out)
}
