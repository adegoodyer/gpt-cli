package main

import "fmt"

type History struct {
	// History is a slice of strings that stores the history of the user's input.
	History []string
}

// NewHistory returns a new History.
func newHistory() *History {
	return &History{}
}

// Add adds a string to the history.
func (h *History) add(s string) {
	h.History = append(h.History, s)
}

// Print prints the history.
func (h *History) print() {
	for i, x := range h.History {
		fmt.Println(i+1, x)
	}
}
