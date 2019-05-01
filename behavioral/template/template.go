package template

// MessageRetriever represents a message which will be returned
// to the user
type MessageRetriever interface {
	Message() string
}

// Template describes the behavior of algorithm to be executed
type Template interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

type TemplateImpl struct{}

func (t *TemplateImpl) first() string {
	return ""
}

func (t *TemplateImpl) third() string {
	return ""
}

func (t *TemplateImpl) ExecuteAlgorithm(m MessageRetriever) string {
	return ""
}
