package instrument

import (
	"testing"

	goast "go/ast"

	"github.com/daspoet/harp/pkg/ast"
	"github.com/daspoet/harp/pkg/pattern"
	"github.com/stretchr/testify/assert"
)

func TestTarget_Accept(t *testing.T) {
	makeDecl := func(typeName, methodName string) goast.FuncDecl {
		return goast.FuncDecl{
			Name: goast.NewIdent(methodName),
			Recv: &goast.FieldList{
				List: []*goast.Field{
					{Type: goast.NewIdent(typeName)},
				},
			},
		}
	}

	var (
		fishSwimPrivate = makeDecl("Fish", "swim")
		fishSwimPublic  = makeDecl("Fish", "Swim")

		duckSwimPrivate = makeDecl("Duck", "swim")
		duckSwimPublic  = makeDecl("Duck", "Swim")
	)

	cases := map[Target]map[goast.FuncDecl]bool{
		{}: {
			fishSwimPrivate: true,
			fishSwimPublic:  true,
			duckSwimPrivate: true,
			duckSwimPublic:  true,
		},
		{Visibility: ast.Public}: {
			fishSwimPrivate: false,
			fishSwimPublic:  true,
			duckSwimPrivate: false,
			duckSwimPublic:  true,
		},
		{Type: pattern.WithWildcards("Duck")}: {
			fishSwimPrivate: false,
			fishSwimPublic:  false,
			duckSwimPrivate: true,
			duckSwimPublic:  true,
		},
		{Method: pattern.WithWildcards("*swim")}: {
			fishSwimPrivate: true,
			fishSwimPublic:  false,
			duckSwimPrivate: true,
			duckSwimPublic:  false,
		},
		{Type: pattern.WithWildcards("Fish"), Visibility: ast.Private}: {
			fishSwimPrivate: true,
			fishSwimPublic:  false,
			duckSwimPrivate: false,
			duckSwimPublic:  false,
		},
	}

	for target, expected := range cases {
		for node, shouldAccept := range expected {
			assert.Equal(t, shouldAccept, target.Accepts(&node), "expected '%#v' to accept '%s'", target, node.Name.Name)
		}
	}
}
