package main

import (
	"strings"
)

type TextGrid struct {
	text [][]rune
}

func NewTextGrid() *TextGrid {
	return &TextGrid{
		text: [][]rune{{}},
	}
}

func NewTextGridFromString(s string) *TextGrid {
	var text [][]rune
	for _, line := range strings.Split(s, "\n") {
		text = append(text, []rune(line))
	}
	return &TextGrid{text}
}

func (tg *TextGrid) String() string {
	var s string
	for i, line := range tg.text {
		if i != 0 {
			s += "\n"
		}
		s += string(line)
	}
	return s
}

func (tg *TextGrid) GetText() [][]rune {
	var text [][]rune
	for _, line := range tg.text {
		text = append(text, append([]rune{}, line...))
	}
	return text
}

func (tg *TextGrid) SetText(text [][]rune) {
	tg.text = text
}

func (tg *TextGrid) NRows() int {
	return len(tg.text)
}

func (tg *TextGrid) MaxCols() int {
	max := 0
	for _, line := range tg.text {
		if n := len(line); n > max {
			max = n
		}
	}
	return max
}

func (tg *TextGrid) MinCols() int {
	min := 0
	for _, line := range tg.text {
		if n := len(line); n < min {
			min = n
		}
	}
	return min
}

func (tg *TextGrid) WidthAt(y int) int {
	return len(tg.text[y]) // TODO: check for out of range
}

func (tg *TextGrid) AddLineAt(y int) {
	before, after := tg.text[:y], tg.text[y:]
	tg.text = append(
		before,
		append([][]rune{{}}, after...)...,
	)
}

func (tg *TextGrid) AddLine() {
	tg.AddLineAt(len(tg.text))
}

func (tg *TextGrid) RemoveLineAt(y int) {
	tg.text = append(
		tg.text[:y],
		tg.text[y+1:]..., // TODO: check for out of range
	)
}

func (tg *TextGrid) AddRuneAt(y, x int, r rune) {
	tg.text[y] = append(
		append(
			tg.text[y][:x],
			r,
		),
		tg.text[y][x:]...,
	)
}

func (tg *TextGrid) AddRuneToEndOfLine(y int, r rune) {
	tg.text[y] = append(tg.text[y], r)
}

func (tg *TextGrid) AddStringAt(y, x int, s string) {
	tg.text[y] = append(
		append(
			tg.text[y][:x],
			[]rune(s)...,
		),
		tg.text[y][x:]...,
	)
}

func (tg *TextGrid) AddStringToEndOfLine(y int, s string) {
	tg.text[y] = append(tg.text[y], []rune(s)...)
}

func (tg *TextGrid) SplitLineAt(y, x int) {
	// Add the new line
	tg.AddLineAt(y + 1)

	// Move the text after `x` to the new line
	tg.text[y+1] = tg.text[y][x:]
	tg.text[y] = tg.text[y][:x]
}
