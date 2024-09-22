package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type statusModel struct {
	spinner  spinner.Model
	progress progress.Model
}

func newStatusModel() statusModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = primaryStyle

	return statusModel{
		spinner:  s,
		progress: progress.New(progress.WithDefaultGradient()),
	}
}

func (m statusModel) Init() tea.Cmd {
	logger.Logf("Init[status]\n")
	return m.spinner.Tick
}

func (m statusModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		logger.Logf("Update[status] | key: `%s`\n", msg.String())
		return m, nil
	}

	m.spinner, cmd = m.spinner.Update(msg)

	return m, cmd
}

func (m statusModel) View() string {
	var b strings.Builder

	b.WriteString(strings.Join([]string{
		m.spinner.View(),
		m.progress.View(),
	}, " ") + "\n")

	return b.String()
}
