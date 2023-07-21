package grammar

import (
	typ "github.com/BananZG/filter-parser/v2/internal/types"
	"github.com/di-wu/parser/ast"
	"github.com/di-wu/parser/op"
)

func Path(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(ast.Capture{
		Type:        typ.Path,
		TypeStrings: typ.Stringer,
		Value: op.Or{
			op.And{
				ValuePath,
				op.Optional(SubAttr),
			},
			AttrPath,
		},
	})
}
