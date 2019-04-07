package linter

import (
	"go/token"
	"testing"
)

func Test_Result_String(t *testing.T) {
	tests := map[string]struct {
		position token.Position
		comment  string
		expected string
	}{
		"Normal": {
			position: token.Position{
				Filename: "file1.go",
				Offset:   0,
				Line:     10,
				Column:   11,
			},
			comment:  "test comment",
			expected: "file1.go:10:11: test comment",
		},
		"Filename is empty": {
			position: token.Position{
				Filename: "",
				Offset:   0,
				Line:     10,
				Column:   11,
			},
			comment:  "test comment",
			expected: "10:11: test comment",
		},
		"Irregular: Line 0": {
			position: token.Position{
				Filename: "file1.go",
				Offset:   0,
				Line:     0, // Normally, this pattern is not happened.
				Column:   11,
			},
			comment:  "test comment",
			expected: "file1.go: test comment",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := NewResult(tt.position, tt.comment)
			output := result.String()

			if tt.expected != output {
				t.Errorf("output does not match. want=%q got=%q", tt.expected, output)
			}
		})
	}
}
