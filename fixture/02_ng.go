package fixture

// Interface02 comment 1st line
// Interface02 comment 2nd line
type Interface02 interface {
	// Inner comment does not check
	Method()
}

// Struct02 comment
type Struct02 struct {
	// this comment is not checked
	Field int64
}

// Method comment
func (s *Struct02) Method() {
	// this comment is not checked
}

// main02 comment
func main02() {
	// comment in main does not check
}
