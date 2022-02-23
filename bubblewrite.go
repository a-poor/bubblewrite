package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type TextArea struct {
	text []string

	w, h int

	showCursor      bool
	showLineNumbers bool
}

func New() *TextArea {
	return &TextArea{}
}

func (ta *TextArea) appendString(s string) {
	ta.text = append(ta.text, s)
}

func (ta *TextArea) popString() {
	if len(ta.text) > 0 {
		ta.text = ta.text[:len(ta.text)-1]
	}
}

func (ta *TextArea) setSize(w, h int) {
	ta.w = w
	ta.h = h
}

func (ta *TextArea) Update(msg tea.Msg) (*TextArea, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch k := msg.String(); {
		case k == "backspace":
			ta.popString()

		case k == "enter":
			ta.appendString("\n")

		case len(k) == 1:
			ta.appendString(k)

		}

	}

	return ta, nil
}

func (ta *TextArea) View() string {
	return ""
}
