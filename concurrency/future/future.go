package future

// SuccessFunc is a type function to be used on success call future
type SuccessFunc func(string)

// FailFunc is a type function to be used on fail
type FailFunc func(error)

//ExecuteStringFunc is a type function to be used on execution
type ExecuteStringFunc func() (string, error)

// MaybeString is a type to use the future pattern
type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

// Success is called when future is ok
func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	s.successFunc = f
	return s
}

// Fail is called when future failed
func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	s.failFunc = f
	return s
}

// Execute executes the function ExecuteStringFunc
func (s *MaybeString) Execute(f ExecuteStringFunc) {
	go func(s *MaybeString) {
		str, err := f()
		if err != nil {
			s.failFunc(err)
		} else {
			s.successFunc(str)
		}
	}(s)
}
