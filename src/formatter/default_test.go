package formatter

import (
	"go/token"
	"testing"

	"github.com/suzuki/go-cmnt-eol-lint/src/linter"
)

func Test_Default_Format(t *testing.T) {
	formatter := NewDefaultFormatter()

	tests := map[string]struct {
		results  []*linter.Result
		expected string
	}{
		"No Results": {
			results:  []*linter.Result{},
			expected: "",
		},
		"Have 1 Result": {
			results: []*linter.Result{
				&linter.Result{
					Position: token.Position{
						Filename: "file1.go",
						Offset:   0,
						Line:     10,
						Column:   11,
					},
					Comment: "test comment",
				},
			},
			expected: "file1.go:10:11: test comment\n",
		},
		"Have 2 Results": {
			results: []*linter.Result{
				&linter.Result{
					Position: token.Position{
						Filename: "file1.go",
						Offset:   0,
						Line:     10,
						Column:   11,
					},
					Comment: "test comment",
				},
				&linter.Result{
					Position: token.Position{
						Filename: "file2.go",
						Offset:   0,
						Line:     20,
						Column:   21,
					},
					Comment: "test comment",
				},
			},
			expected: "file1.go:10:11: test comment\nfile2.go:20:21: test comment\n",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			output, err := formatter.Format(tt.results)
			if err != nil {
				t.Fatalf("err is not nil. got=%q", err)
			}
			if tt.expected != output {
				t.Errorf("output is not match. want=%q got=%q", tt.expected, output)
			}
		})
	}
}
