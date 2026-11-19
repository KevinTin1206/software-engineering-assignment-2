package main

import (
	"fmt"
	"strings"
)

// Price liefert den Preis in Cent. Display liefert die Darstellung des Elements.
type Item interface {
	Price() int
	Display() string
}

type Glass struct {
	cents int
}

func (g Glass) Price() int      { return g.cents }
func (g Glass) Display() string { return "Glas" }

type Frame struct {
	cents int
}

func (f Frame) Price() int      { return f.cents }
func (f Frame) Display() string { return "Rahmen" }

type Window struct {
	name  string
	items []Item
}

// Das erste Element ist Pflicht, damit der Konstruktor die Kardinalität 1..* abbildet.
func NewWindow(name string, first Item, rest ...Item) Window {
	return Window{name: name, items: append([]Item{first}, rest...)}
}

func (w Window) Price() int {
	total := 0
	for _, item := range w.items {
		total += item.Price()
	}
	return total
}

func (w Window) Display() string {
	var result strings.Builder
	result.WriteString(w.name)
	for _, item := range w.items {
		result.WriteString("\n  ")
		result.WriteString(strings.ReplaceAll(item.Display(), "\n", "\n  "))
	}
	return result.String()
}

func main() {
	subwindow := NewWindow("Teilfenster",
		Frame{cents: 1000},
		Glass{cents: 500},
		Glass{cents: 500},
		Glass{cents: 500},
		Glass{cents: 500},
	)
	var window Item = NewWindow("Fenster",
		Frame{cents: 3000},
		Glass{cents: 1000},
		Glass{cents: 1000},
		Glass{cents: 1000},
		subwindow,
	)

	fmt.Println(window.Display())
	price := window.Price()
	fmt.Printf("Gesamtpreis: %d,%02d €\n", price/100, price%100)
}
