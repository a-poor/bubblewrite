package textgrid

import "strings"

// gridify converts a string into a 2D rune slice where the first
// axis is the lines of the string, and the second axis is the
// runes in each line.
//
// For example the string "hello\nworld" would be converted to:
//
// [][]rune{[]rune("hello"), []rune("world")}
//
func gridify(s string) [][]rune {
	var grid [][]rune
	for _, line := range strings.Split(s, "\n") {
		grid = append(grid, []rune(line))
	}
	return grid
}

// ungridify converts a slice of slices of runes into a single string.
// The first axis is the lines of the string, and the second axis is
// the runes in each line.
//
// For example the slice [][]rune{[]rune("hello"), []rune("world")}
// would be converted to "hello\nworld".
func ungridify(g [][]rune) string {
	var res string
	for i, line := range g {
		if i != 0 {
			res += "\n"
		}
		res += string(line)
	}
	return res
}

// duplicate returns a copy of the rune slice where the underlying
// data isn't shared with the original.
func duplicate(s []rune) []rune {
	return append([]rune{}, s...)
}

// duplicate2D returns a copy of the slice of rune slices where the underlying
// data isn't shared with the original.
func duplicate2D(s [][]rune) [][]rune {
	var res [][]rune
	for _, line := range s {
		res = append(res, duplicate(line))
	}
	return res
}

// Join together multiple rune slices into a single rune slice
func join(left []rune, right ...[]rune) []rune {
	res := duplicate(left)
	for _, r := range right {
		res = append(res, r...)
	}
	return res
}

// Join together multiple rune slices into a single rune slice
func join2D(left [][]rune, right ...[][]rune) [][]rune {
	res := duplicate2D(left)
	for _, r := range right {
		res = append(res, duplicate2D(r)...)
	}
	return res
}

// split returns two rune slices, the first containing items in s
// up to the index at, and the second containing items after.
//
// If at is out of bounds, one of the slices will be empty
// (if at <= 0, the left slice will be empty, and if at >= len(s),
// the right slice will be empty).
//
// Additionally, split calls duplicate on the returned slices, so that
// the underlying data of the returned slices aren't shared with the original.
func split(s []rune, at int) ([]rune, []rune) {
	if at <= 0 {
		return []rune{}, duplicate(s)
	}
	if at >= len(s) {
		return duplicate(s), []rune{}
	}
	return duplicate(s[:at]), duplicate(s[at:])
}

// split2D returns two slices of rune slices, the first containing items in s
// up to the index at, and the second containing items after.
//
// If at is out of bounds, one of the slices will be empty
// (if at <= 0, the left slice will be empty, and if at >= len(s),
// the right slice will be empty).
//
// Additionally, split calls duplicate on the returned slices, so that
// the underlying data of the returned slices aren't shared with the original.
func split2D(s [][]rune, at int) ([][]rune, [][]rune) {
	if at <= 0 {
		return [][]rune{}, duplicate2D(s)
	}
	if at >= len(s) {
		return duplicate2D(s), [][]rune{}
	}
	return duplicate2D(s[:at]), duplicate2D(s[at:])
}

// insertAt inserts rune r into the rune slice s at index at.
// This is a 1D version of insertAt2D.
func insertAt(s []rune, at int, r rune) []rune {
	left, right := split(s, at)
	return join(append(left, r), right)
}

// insertAt2D inserts rune slice, r, into the slice of rune slices, s, at index, at.
// This is a 2D version of insertAt.
func insertAt2D(s [][]rune, at int, r []rune) [][]rune {
	left, right := split2D(s, at)
	return join2D(append(left, r), right)
}
