package shapes

import (
	"fmt"
	"os"

	"github.com/leogsouza/go-design-patterns/behavioral/strategy"
)

const (
	TextStrategy  = "text"
	ImageStrategy = "image"
)

func NewPrinter(s string) (strategy.Output, error) {
	switch s {
	case TextStrategy:
		return &TextSquare{
			PrintOutput: strategy.PrintOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	case ImageStrategy:
		return &ImageSquare{
			PrintOutput: strategy.PrintOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found\n", s)
	}
}
