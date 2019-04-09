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
			expected: "file1.go:9:11: test comment",
		},
		"Filename is empty": {
			position: token.Position{
				Filename: "",
				Offset:   0,
				Line:     10,
				Column:   11,
			},
			comment:  "test comment",
			expected: "9:11: test comment",
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

func Test_GetPosition(t *testing.T) {
	tests := map[string]struct {
		position token.Position
		comment  string
		expected string
	}{
		"Expected output is line 9": {
			position: token.Position{
				Filename: "file1.go",
				Offset:   0,
				Line:     10, // `Line` should be decremented in output.
				Column:   11,
			},
			comment:  "test comment",
			expected: "file1.go:9:11",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := NewResult(tt.position, tt.comment)
			if tt.expected != result.GetPosition() {
				t.Errorf("GetPosition() does not match. want=%q got=%q", tt.expected, result.GetPosition())
			}
		})
	}
}

func Test_GetComment(t *testing.T) {
	tests := map[string]struct {
		position token.Position
		comment  string
		expected string
	}{
		"replace new line to space": {
			position: token.Position{},
			comment:  "test comment 1\ntest comment 2",
			expected: "test comment 1 test comment 2", // "\n" replaced to " "
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := NewResult(tt.position, tt.comment)
			if tt.expected != result.GetComment() {
				t.Errorf("GetComment() does not match. want=%q got=%q", tt.expected, result.GetComment())
			}
		})
	}
}
