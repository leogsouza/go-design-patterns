package adapter

import "fmt"

// LegacyPrinter is an old interface
type LegacyPrinter interface {
	Print(s string) string
}

// MyLegacyPrinter implements LegacyPrinter
type MyLegacyPrinter struct{}

// Print accepts a string and return a message
func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	fmt.Println(newMsg)
	return
}

// ModernPrinter is an new interface
type ModernPrinter interface {
	PrintStored() string
}

// PrinterAdapter works as an adapter from
// old interface to new interface
type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

// PrintStored prints the message store in legacy printer
func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}
	return
}
