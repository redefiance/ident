package ident

import (
	"path"
	"path/filepath"

	"github.com/npat-efault/godef/exp-go/ast"
	"github.com/npat-efault/godef/exp-go/parser"
	"github.com/npat-efault/godef/exp-go/token"
	"github.com/npat-efault/godef/exp-go/types"
)

var fileset = types.FileSet
var scopes = map[string]*ast.Scope{}

func getScope(filepath string) *ast.Scope {
	dirpath := path.Base(filepath)
	scope, ok := scopes[dirpath]
	if !ok {
		scope = ast.NewScope(parser.Universe)
		scopes[dirpath] = scope
	}
	return scope
}

func getDefPosition(expr ast.Expr) *token.Position {
	obj, _ := types.ExprType(expr, types.DefaultImporter)
	if obj == nil {
		return nil
	}
	pos := fileset.Position(types.DeclPos(obj))
	if realname, err := filepath.EvalSymlinks(pos.Filename); err == nil {
		pos.Filename = realname
	}
	return &pos
}
