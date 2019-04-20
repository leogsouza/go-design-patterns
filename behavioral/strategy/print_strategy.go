package strategy

import "io"

type Output interface {
	Print() error
	SetLog(io.Writer)
	SetWriter(io.Writer)
}

type PrintOutput struct {
	Writer    io.Writer
	LogWriter io.Writer
}

func (d *PrintOutput) SetLog(w io.Writer) {
	d.LogWriter = w
}

func (d *PrintOutput) SetWriter(w io.Writer) {
	d.Writer = w
}
