package shapes

import (
	"github.com/leogsouza/go-design-patterns/behavioral/strategy"
)

type TextSquare struct {
	strategy.PrintOutput
}

func (t *TextSquare) Print() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}
