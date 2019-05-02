package template

import "strings"

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
	return "hello"
}

func (t *Template) third() string {
	return "template"
}

// ExecuteAlgorithm accepts MessageRetriever as argument and returns
// the full algorithm: a single string done by joining the strings returned
// by the first(), Message() string and third() methods
func (t *Template) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}
