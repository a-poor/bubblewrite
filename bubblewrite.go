package main

import (
	"fmt"
	"strings"

	"github.com/a-poor/bubblewrite/textgrid"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	if ta.cursor.x == ta.text.WidthAt(ta.cursor.y) && ta.cursor.y == ta.text.NRows() {
		return
	}

	// If at the end of a (non-last) line, move down a line
	if ta.cursor.x == ta.text.WidthAt(ta.cursor.y) {
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
	if ta.cursor.x == ta.text.WidthAt(ta.cursor.y) {
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
	if ta.cursor.y == ta.text.NRows() {
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

	nextX := ta.cursor.x - 1
	if ta.cursor.x == 0 {
		nextX = ta.text.WidthAt(ta.cursor.y - 1)
	}

	ta.text.DeleteRuneAt(ta.cursor.y, ta.cursor.x)
	ta.cursor.x = nextX
	ta.cursor.y--
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
	// Create the line-numbers side-bar
	var nums []string
	for i := 0; i < ta.h; i++ {
		nums = append(
			nums,
			fmt.Sprintf(
				"%d",
				i+1+ta.offset,
			),
		)
	}
	lineNums := strings.Join(nums, "\n")
	lineNums = lineNumberStyle.Render(lineNums)

	// Format the result
	block := ta.text.String()

	// Join them and return
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lineNums,
		block,
	)
}
