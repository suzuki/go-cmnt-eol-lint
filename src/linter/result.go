package linter

import (
	"fmt"
	"go/token"
	"strings"
)

// Result is struct
type Result struct {
	Position token.Position
	Comment  string
}

func NewResult(position token.Position, comment string) *Result {
	// The `position` points out `type`, `func`, etc line. So, `position.Line` decrements here.
	position.Line = position.Line - 1

	return &Result{
		Position: position,
		Comment:  comment,
	}
}

func (r *Result) String() string {
	return fmt.Sprintf("%s: %s", r.GetPosition(), r.GetComment())
}

func (r *Result) GetPosition() string {
	return fmt.Sprintf("%s", r.Position.String())
}

func (r *Result) GetComment() string {
	return strings.ReplaceAll(r.Comment, "\n", " ")
}
