package config

import (
	"github.com/whitepages/terraform-provider-dummy/Godeps/_workspace/src/github.com/hashicorp/terraform/config/lang/ast"
)

type noopNode struct{}

func (n *noopNode) Accept(ast.Visitor) ast.Node      { return n }
func (n *noopNode) Pos() ast.Pos                     { return ast.Pos{} }
func (n *noopNode) Type(ast.Scope) (ast.Type, error) { return ast.TypeString, nil }
