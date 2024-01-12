// Copyright 2023 The GoPlus Authors (goplus.org). All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parserutil

import (
	"github.com/goplus/gop/ast"
	"github.com/goplus/gop/parser"
	"github.com/goplus/gop/token"
)

const (
	// ParseHeader specifies that the main package declaration and imports are needed.
	// This is the mode used when attempting to examine the package graph structure.
	ParseHeader = parser.AllErrors | parser.ParseComments | parser.ImportsOnly

	// ParseFull specifies the full AST is needed.
	// This is used for files of direct interest where the entire contents must
	// be considered.
	ParseFull = parser.AllErrors | parser.ParseComments

	// SkipObjectResolution - don't resolve identifiers to objects - see ParseFile
	SkipObjectResolution = parser.SkipObjectResolution
)

func ParseFile(fset *token.FileSet, filename string, src interface{}, mode parser.Mode) (f *ast.File, err error) {
	if filename != "" {
		f, err = parser.ParseEntry(fset, filename, src, parser.Config{Mode: mode})
	} else {
		f, err = parser.ParseFile(fset, filename, src, mode)
	}
	return
}
