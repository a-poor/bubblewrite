package main

type TextGrid struct {
	text [][]rune
}

func NewTextGrid() *TextGrid {
	return &TextGrid{
		text: [][]rune{{}},
	}
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
	tg.text = append(
		append(
			tg.text[:y],
			[]rune{},
		),
		tg.text[y:]...,
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
	tg.text[r] = append(
		append(
			tg.text[y][:x],
			r,
		),
		tg.text[y][x:]...,
	)
}
