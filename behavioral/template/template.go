package template

// MessageRetriever represents a message which will be returned
// to the user
type MessageRetriever interface {
	Message() string
}

// Templater describes the behavior of algorithm to be executed
type Templater interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

// Template represents the template interface implementation
type Template struct{}

func (t *Template) first() string {
	return ""
}

func (t *Template) third() string {
	return ""
}

// ExecuteAlgorithm executes the altorithm and returns the message
func (t *Template) ExecuteAlgorithm(m MessageRetriever) string {
	return ""
}
