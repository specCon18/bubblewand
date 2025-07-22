package tui

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	count int
}

func (m model) Init() tea.Cmd {
	return nil
}
