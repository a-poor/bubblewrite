package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TextArea struct {
	text [][]rune

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
		text: [][]rune{[]rune{}},
	}
}

func (ta *TextArea) moveCursorRight() {
	//...
}

func (ta *TextArea) moveCursorLeft() {
	//...
}

func (ta *TextArea) moveCursorUp() {
	//...
}

func (ta *TextArea) moveCursorDown() {
	//...
}

func (ta *TextArea) insertRuneAtCursor(r rune) {
	ta.text[ta.cursor.y] = append(
		append(
			ta.text[ta.cursor.y][:ta.cursor.x],
			r,
		),
		ta.text[ta.cursor.y][ta.cursor.x:]...,
	)
	ta.cursor.x++
}

func (ta *TextArea) deleteRuneAtCursor() {
	// If at (0, 0), do nothing
	if ta.cursor.x == 0 && ta.cursor.y == 0 {
		return
	}

	// If not at the beginning of a line...
	if ta.cursor.x != 0 {
		// ...delete the rune
		ta.text[ta.cursor.y] = append(
			ta.text[ta.cursor.y][:ta.cursor.x],
			ta.text[ta.cursor.y][ta.cursor.x:]...,
		)
		ta.cursor.x--
		return
	}

	// Otherwise, merge it with the previous line
	newX := len(ta.text[ta.cursor.y-1])
	ta.text[ta.cursor.y-1] = append(
		ta.text[ta.cursor.y-1],
		ta.text[ta.cursor.y]...,
	)
	ta.text = append(
		ta.text[:ta.cursor.y],
		ta.text[ta.cursor.y+1:]...,
	)
	ta.cursor.y--
	ta.cursor.x = newX
}

func (ta *TextArea) insertNewlineAtCursor() {
	// Part of the current line staying in place
	lstay := ta.text[ta.cursor.y][:ta.cursor.x]

	// Part of the current line moving to the new line
	lgo := ta.text[ta.cursor.y][ta.cursor.x:]

	// Update the text grid
	ta.text = append(
		append(
			ta.text[:ta.cursor.y],
			lstay,
		),
		lgo,
	)

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
			ta.moveCursorDown()

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
	var block string
	for i, line := range ta.text {
		if i != 0 {
			block += "\n"
		}
		block += string(line)
	}

	// Join them and return
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lineNums,
		block,
	)
}
