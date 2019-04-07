package formatter

import (
	"strings"

	"github.com/suzuki/go-cmnt-eol-lint/src/linter"
)

type Formatter interface {
	Format([]*linter.Result) (string, error)
}

type defaultFormatter struct{}

func NewDefaultFormatter() Formatter {
	return &defaultFormatter{}
}

// Format convert to the `errorformat` text from the Results.
func (f *defaultFormatter) Format(results []*linter.Result) (string, error) {
	var output string

	for _, r := range results {
		output += strings.ReplaceAll(r.String(), "\n", " ")
		output += "\n"
	}

	return output, nil
}
