package main

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	ta *TextArea
}

func newModel() *Model {
	return &Model{New()}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch k := msg.String(); {
		case k == "ctrl+c":
			return m, tea.Quit
		}
	}

	m.ta.Update(msg)

	return m, nil
}

func (m Model) View() string {
	return m.ta.View()
}

func main() {
	m := newModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	p.Start()
}
