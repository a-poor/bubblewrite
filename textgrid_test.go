package main_test

import (
	"testing"

	. "github.com/a-poor/bubblewrite"
)

func TestCreateTextGrid(t *testing.T) {
	// Confirm that at least this works
	_ = NewTextGrid()
}

func TestAddRows(t *testing.T) {
	// Create a TG
	tg := NewTextGrid()

	// It should already have one row
	if n := tg.NRows(); n != 1 {
		t.Errorf("Expected 1, got %d", n)
	}

	// Add 3 more rows and check
	tg.AddLine()
	tg.AddLine()
	tg.AddLine()
	if n := tg.NRows(); n != 4 {
		t.Errorf("Expected 3 rows, got %d", n)
	}

	// Remove some rows and check
	tg.RemoveLineAt(1)
	tg.RemoveLineAt(1)
	if n := tg.NRows(); n != 2 {
		t.Errorf("Expected 2 rows, got %d", n)
	}
}

func TestGetSetText(t *testing.T) {
	tg := NewTextGrid()
	if s := tg.String(); s != "" {
		t.Errorf("Expected empty string, got %s", s)
	}

	tg.SetText([][]rune{[]rune("Hello"), []rune("World")})
	if s := tg.String(); s != "Hello\nWorld" {
		t.Errorf("Expected 'Hello\nWorld', got %s", s)
	}

	if s := tg.GetText()[0][0]; s != 'H' {
		t.Errorf("Expected 'H', got %c", s)
	}
}

func TestAddRune(t *testing.T) {
	tg := NewTextGrid()
	for i, r := range []rune("Hello") {
		tg.AddRuneAt(0, i, r)
	}
	if s := tg.String(); s != "Hello" {
		t.Errorf("Expected 'Hello', got %s", s)
	}
}

func TestAddLineInMiddle(t *testing.T) {
	// Create a TG with two lines
	tg := NewTextGrid()
	tg.AddLine()

	// Write "hello world" on two lines
	tg.AddStringToEndOfLine(0, "Hello")
	tg.AddStringToEndOfLine(1, "World")

	// Check that the text is correct
	if s := tg.String(); s != "Hello\nWorld" {
		t.Errorf("Expected %q, got %q", "Hello\nWorld", s)
		t.FailNow()
	}

	// Add a new line in the middle
	tg.AddLineAt(1)

	// Check that the text is correct
	if s := tg.String(); s != "Hello\n\nWorld" {
		t.Errorf("Expected %q, got %q", "Hello\n\nWorld", s)
	}
}

func TestSplitLineAt(t *testing.T) {
	// Create a TG
	tg := NewTextGrid()

	// Write the starter string (and check)
	tg.AddStringAt(0, 0, "HelloWorld")
	if s := tg.String(); s != "HelloWorld" {
		t.Errorf("Expected %q, got %q", "HelloWorld", s)
	}

	// Split the line at the middle (and check)
	tg.SplitLineAt(0, 5)
	if s := tg.String(); s != "Hello\nWorld" {
		t.Errorf("Expected %q, got %q", "Hello\nWorld", s)
	}
}
