package fixture

// Interface01 comment 1st line
// Interface01 comment 2nd line.
type Interface01 interface {
	// Inner comment does not check
	Method()
}

// Struct01 comment!
type Struct01 struct {
	// this comment is not checked
	Field int64
}

// Method (comment)
func (s *Struct01) Method() {
	// this comment is not checked
}

// main01 [comment]
func main01() {
	// comment in main does not check
}
