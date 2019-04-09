package linter

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"

	"github.com/suzuki/go-cmnt-eol-lint/src/env"
)

type Linter interface {
	Lint(src string) ([]*Result, error)
	LintFile(path string) ([]*Result, error)
}

type defaultLinter struct {
	config *env.Config
}

func New(config *env.Config) Linter {
	return &defaultLinter{
		config: config,
	}
}

func (l *defaultLinter) Lint(src string) ([]*Result, error) {
	return l.lint("", []byte(src))
}

func (l *defaultLinter) LintFile(path string) ([]*Result, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return []*Result{}, nil
	}

	src, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return l.lint(path, src)
}

func (l *defaultLinter) hasAllowedEOL(comment string) bool {
	r := []rune(strings.TrimSpace(comment))
	size := len(r)

	if size < 1 {
		return true
	}

	allowedEOLs := []rune(l.config.AllowedEOLs)
	eol := r[size-1]

	for _, allow := range allowedEOLs {
		if eol == allow {
			return true
		}
	}

	return false
}

func (l *defaultLinter) hasMinWords(comment string) bool {
	return len(strings.Fields(comment)) >= l.config.MinWords
}

func (l *defaultLinter) lint(path string, src []byte) ([]*Result, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var results []*Result

	for _, decl := range f.Decls {
		switch v := decl.(type) {
		case *ast.GenDecl:
			comment := v.Doc.Text()
			if !l.hasMinWords(comment) {
				continue
			}

			// GenDecl means `import`, `const`, `type`, `var`
			if !l.hasAllowedEOL(comment) {
				r := NewResult(fset.Position(v.Pos()), comment)
				results = append(results, r)
			}

		case *ast.FuncDecl:
			comment := v.Doc.Text()
			if !l.hasMinWords(comment) {
				continue
			}

			// FuncDecl means `func`
			if !l.hasAllowedEOL(comment) {
				r := NewResult(fset.Position(v.Pos()), comment)
				results = append(results, r)
			}
		}
	}

	return results, nil
}
