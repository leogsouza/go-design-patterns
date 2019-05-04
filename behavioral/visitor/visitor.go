package visitor

import (
	"fmt"
	"io"
	"os"
)

// MessageA is a message logger that will print a message
type MessageA struct {
	Msg    string
	Output io.Writer
}

// MessageB is a message logger that will print a message
type MessageB struct {
	Msg    string
	Output io.Writer
}

// Visitor has two methods that able to modify the message to be printed
// using your Visit methods
type Visitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}

// Visitable accepts the visitor that will execute the decoupled algorithm
type Visitable interface {
	Accept(Visitor)
}

// MessageVisitor implements Visitor and Visitable interfaceds
type MessageVisitor struct{}

// VisitA stores the modified message into messageA
func (mv *MessageVisitor) VisitA(m *MessageA) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited A)")
}

// VisitB stores the modified message into messageB
func (mv *MessageVisitor) VisitB(m *MessageB) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited B)")
}

// Accept add messageA to VisitA
func (m *MessageA) Accept(v Visitor) {
	v.VisitA(m)
}

// Accept add messageB to VisitB
func (m *MessageB) Accept(v Visitor) {
	v.VisitB(m)
}

// Print prints the MessageA message
func (m *MessageA) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}

	fmt.Fprintf(m.Output, "A: %s", m.Msg)
}

// Print prints the MessageB message
func (m *MessageB) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}

	fmt.Fprintf(m.Output, "B: %s", m.Msg)
}
