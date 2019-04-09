package fixture

// Interface02 comment 1st line
// Interface02 comment 2nd line
type Interface02 interface {
	// Inner comment does not check
	Method()
}

// Struct02 is a structure for checking the comment eol
type Struct02 struct {
	// this comment is not checked
	Field int64
}

// Method is a method for checking the comment eol
func (s *Struct02) Method() {
	// this comment is not checked
}

// main02 is a main func for checking the comment eol
func main02() {
	// comment in main02 does not check
}

// short-comment no-check
func func02() {
	// comment in func02 does not check
}
