package main

import "github.com/charmbracelet/lipgloss"

var lineNumberStyle = lipgloss.NewStyle().
	Width(3).
	Background(lipgloss.Color("#faa")).
	Foreground(lipgloss.Color("#333")).
	Align(lipgloss.Right).
	MarginRight(1)

var regularStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#000")).
	Foreground(lipgloss.Color("#ddd"))

var activeLineStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#ddd")).
	Foreground(lipgloss.Color("#333"))

var cursorStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#fff")).
	Foreground(lipgloss.Color("#444"))
