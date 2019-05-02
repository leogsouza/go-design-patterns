package template

import (
	"strings"
	"testing"
)

type TestStruct struct {
	Templater
}

func (m *TestStruct) Message() string {
	return "world"
}

func TestTemplate_ExecuteAlgorithm(t *testing.T) {
	t.Run("Using interfaces", func(t *testing.T) {
		s := &TestStruct{Templater: &Template{}}
		res := s.ExecuteAlgorithm(s)
		expected := "world"

		if !strings.Contains(res, expected) {
			t.Errorf("Expected string '%s' wasn't found on retuerned string\n",
				expected)
		}
	})
}
