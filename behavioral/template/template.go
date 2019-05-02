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

// Template represents the templater interface implementation
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

// AnonymousTemplate represents the templater interface implementation
// using anonymous functions
type AnonymousTemplate struct{}

func (a *AnonymousTemplate) first() string {
	return "hello"
}

func (a *AnonymousTemplate) third() string {
	return "template"
}

// ExecuteAlgorithm accepts MessageRetriever as argument and returns
// the full algorithm: a single string done by joining the strings returned
// by the first(), Message() string and third() methods
func (a *AnonymousTemplate) ExecuteAlgorithm(f func() string) string {
	return strings.Join([]string{a.first(), f(), a.third()}, " ")
}

// adapter represents the templater interface implementation
// using anonymous functions and adapter
type adapter struct {
	myFunc func() string
}

// Message returns a message from anonymous function
func (a *adapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}

	return ""
}

// MessageRetrieverAdapter returns a MessageRetriever type
func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &adapter{myFunc: f}
}
