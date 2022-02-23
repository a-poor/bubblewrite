package main

import "github.com/charmbracelet/lipgloss"

var lineNumberStyle = lipgloss.NewStyle().
	Width(3).
	Background(lipgloss.Color("#faa")).
	Foreground(lipgloss.Color("#333")).
	Align(lipgloss.Right).
	MarginRight(1)
