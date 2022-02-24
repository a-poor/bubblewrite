package main

import (
	"fmt"

	"github.com/a-poor/bubblewrite/textgrid"
	tea "github.com/charmbracelet/bubbletea"
)

type TextArea struct {
	text *textgrid.TextGrid

	w, h   int
	offset int
	cursor struct {
		show bool
		x, y int
	}
	showLineNumbers bool
}

func New() *TextArea {
	return &TextArea{
		text: textgrid.NewTextGrid(),
	}
}

func (ta *TextArea) moveCursorRight() {
	// If all the way to the right on the last line, do nothing
	if ta.cursor.x == ta.text.WidthAt(ta.cursor.y) && ta.cursor.y >= ta.text.NRows()-1 {
		return
	}

	// If at the end of a (non-last) line, move down a line
	if ta.cursor.x >= ta.text.WidthAt(ta.cursor.y)-1 {
		ta.cursor.y++
		ta.cursor.x = 0
		return
	}

	// Otherwise, move right
	ta.cursor.x++
}

func (ta *TextArea) moveCursorLeft() {
	// If all the way to the left on the first line, do nothing
	if ta.cursor.x == 0 && ta.cursor.y == 0 {
		return
	}

	// If at the start of a (non-first) line, move up a line
	if ta.cursor.x == 0 {
		ta.cursor.y--
		ta.cursor.x = ta.text.WidthAt(ta.cursor.y)
		return
	}

	// Otherwise, move left
	ta.cursor.x--
}

func (ta *TextArea) moveCursorUp() {
	if ta.cursor.y == 0 {
		ta.cursor.x = 0
		return
	}

	ta.cursor.y--
	if w := ta.text.WidthAt(ta.cursor.y); ta.cursor.x > w {
		ta.cursor.x = w
	}
}

func (ta *TextArea) moveCursorDown() {
	if ta.cursor.y >= ta.text.NRows()-1 {
		ta.cursor.x = ta.text.WidthAt(ta.cursor.y)
		return
	}

	ta.cursor.y++
	if w := ta.text.WidthAt(ta.cursor.y); ta.cursor.x > w {
		ta.cursor.x = w
	}
}

func (ta *TextArea) insertRuneAtCursor(r rune) {
	ta.text.AddRuneAt(ta.cursor.y, ta.cursor.x, r)
	ta.cursor.x++
}

func (ta *TextArea) deleteRuneAtCursor() {
	if ta.cursor.x == 0 && ta.cursor.y == 0 {
		return
	}

	var movedLine bool
	nextX := ta.cursor.x - 1
	if ta.cursor.x == 0 {
		nextX = ta.text.WidthAt(ta.cursor.y - 1)
		movedLine = true
	}

	ta.text.DeleteRuneAt(ta.cursor.y, ta.cursor.x)
	ta.cursor.x = nextX
	if movedLine {
		ta.cursor.y--
	}
}

func (ta *TextArea) insertNewlineAtCursor() {
	// Split the current line at the cursor
	ta.text.SplitLineAt(ta.cursor.y, ta.cursor.x)

	// Update the cursor positions
	ta.cursor.y++
	ta.cursor.x = 0
}

func (ta *TextArea) setSize(w, h int) {
	ta.w = w
	ta.h = h
}

func (ta *TextArea) Update(msg tea.Msg) (*TextArea, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch k := msg.String(); {
		case k == "ctrl+c":
			return ta, tea.Quit

		case k == "backspace":
			ta.deleteRuneAtCursor()

		case k == "enter":
			ta.insertNewlineAtCursor()

		case k == "up":
			ta.moveCursorUp()

		case k == "down":
			ta.moveCursorDown()

		case k == "left":
			ta.moveCursorLeft()

		case k == "right":
			ta.moveCursorRight()

		case len(k) == 1:
			ta.insertRuneAtCursor([]rune(k)[0])
		}

	case tea.WindowSizeMsg:
		ta.setSize(msg.Width, msg.Height)

	}

	return ta, nil
}

func (ta *TextArea) View() string {
	// Get the raw rune text
	txt := ta.text.GetText()

	// Create a var to store the output
	var block string
	for i, line := range txt {
		if i != 0 {
			block += "\n"
		}
		s := string(line)

		if ta.cursor.y != i {
			block += regularStyle.Render(s)
			continue
		}

		// Otherwise, it's the active line
		s = ""
		for j, r := range append(line, ' ') {
			if j == ta.cursor.x {
				s += cursorStyle.Render(string(r))
			} else {
				s += activeLineStyle.Render(string(r))
			}
		}
		block += s
	}

	// Join them and return
	return fmt.Sprintf(
		"Cursor: %+v\n%s",
		ta.cursor,
		block,
	)
}
