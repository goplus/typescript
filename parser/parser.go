package parser

import (
	"github.com/microsoft/typescript-go/ast"
	"github.com/microsoft/typescript-go/core"

	"github.com/microsoft/typescript-go/internal/parser"
)

func ParseSourceFile(opts ast.SourceFileParseOptions, sourceText string, scriptKind core.ScriptKind) *ast.SourceFile {
	return parser.ParseSourceFile(opts, sourceText, scriptKind)
}
