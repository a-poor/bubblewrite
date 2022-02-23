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
