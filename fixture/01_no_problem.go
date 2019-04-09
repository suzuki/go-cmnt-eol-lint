package fixture

// Interface01 comment 1st line
// Interface01 comment 2nd line.
type Interface01 interface {
	// Inner comment does not check
	Method()
}

// Struct01 is a structure for checking the comment eol!
type Struct01 struct {
	// this comment is not checked
	Field int64
}

// Method is a method for checking the comment (eol)
func (s *Struct01) Method() {
	// this comment is not checked
}

// main01 is a main func for checking the comment [eol]
func main01() {
	// comment in main01 does not check
}
