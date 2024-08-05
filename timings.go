package harp

import (
	"fmt"
	goast "go/ast"
	"go/token"

	"github.com/daspoet/harp/pkg/ast"
	"github.com/daspoet/harp/pkg/instrument"
)

type Timings struct {
	Targets []instrument.Target
}

func (t Timings) Instrument(ctx *instrument.Context, node goast.Node) goast.Node {
	decl, ok := node.(*goast.FuncDecl)
	if !ok {
		return nil
	}
	if t.ignored(decl) {
		return nil
	}

	t.addImports(ctx)
	return t.instrumentMethod(decl)
}

func (t Timings) ignored(node *goast.FuncDecl) bool {
	for _, target := range t.Targets {
		if target.Accepts(node) {
			return false
		}
	}
	return true
}

func (t Timings) addImports(ctx *instrument.Context) {
	ctx.AddImports("fmt", "time")
}

func (t Timings) instrumentMethod(node *goast.FuncDecl) goast.Node {
	beforeAssign := &goast.AssignStmt{
		Lhs: []goast.Expr{
			&goast.Ident{Name: "_before"},
		},
		Rhs: []goast.Expr{
			&goast.CallExpr{
				Fun: &goast.SelectorExpr{
					X:   &goast.Ident{Name: "time"},
					Sel: &goast.Ident{Name: "Now"},
				},
			},
		},
		Tok: token.DEFINE,
	}

	endCall := &goast.DeferStmt{
		Call: &goast.CallExpr{
			Fun: &goast.FuncLit{
                Type: &goast.FuncType{},
				Body: &goast.BlockStmt{
					List: []goast.Stmt{
						&goast.ExprStmt{
							X: &goast.CallExpr{
								Fun: &goast.SelectorExpr{
									X:   &goast.Ident{Name: "fmt"},
									Sel: &goast.Ident{Name: "Println"},
								},
								Args: []goast.Expr{
									&goast.BasicLit{
										Kind:  token.STRING,
										Value: fmt.Sprintf(`"%s.%s took"`, *ast.ReceiverTypeName(node), node.Name.Name),
									},
									&goast.CallExpr{
										Fun: &goast.SelectorExpr{
											X:   &goast.Ident{Name: "time"},
											Sel: &goast.Ident{Name: "Since"},
										},
										Args: []goast.Expr{
											&goast.Ident{Name: "_before"},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	node.Body.List = append([]goast.Stmt{beforeAssign, endCall}, node.Body.List...)
	return node
}

var _ instrument.Instrumentor = (*Timings)(nil)
