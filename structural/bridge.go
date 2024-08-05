package structural

import "fmt"

type Device interface {
	Print()
	SetPrinter(Printer)
}

type Printer interface {
	PrintFile()
}

type MacComputer struct {
	printer Printer
}

func (m *MacComputer) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *MacComputer) SetPrinter(p Printer) {
	m.printer = p
}

type WinComputer struct {
	printer Printer
}

func (w *WinComputer) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *WinComputer) SetPrinter(p Printer) {
	w.printer = p
}

type Epson struct {
}

func (e *Epson) PrintFile() {
	fmt.Println("Printing by a Epson Printer")
}

type Hp struct {
}

func (h *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

func DoBridge() {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &MacComputer{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &WinComputer{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
