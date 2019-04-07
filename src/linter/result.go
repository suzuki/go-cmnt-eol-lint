package linter

import (
	"fmt"
	"go/token"
)

// Result is struct
type Result struct {
	Position token.Position
	Comment  string
}

func NewResult(position token.Position, comment string) *Result {
	return &Result{
		Position: position,
		Comment:  comment,
	}
}

func (r *Result) String() string {
	return fmt.Sprintf("%s: %s", r.Position.String(), r.Comment)
}
