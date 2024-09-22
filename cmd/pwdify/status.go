package main

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type statusTickMsg struct{}

type statusModel struct {
	progress progress.Model
	state    *state
}

func newStatusModel(s *state) statusModel {
	return statusModel{
		progress: progress.New(progress.WithDefaultGradient()),
		state:    s,
	}
}

func (m statusModel) Init() tea.Cmd {
	return m.encrypt
}

func (m statusModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		logger.Logf("Update[status] | key: `%s`\n", msg.String())
		return m, nil

	case statusTickMsg:
		logger.Logf("Update[status] | tick\n")
		return m, m.tick()
	}

	return m, cmd
}

func (m statusModel) View() string {
	var b strings.Builder

	b.WriteString(m.progress.View() + "\n")

	return b.String()
}

func (m statusModel) encrypt() tea.Msg {
	return statusTickMsg{}
}

func (m statusModel) tick() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return statusTickMsg{}
	})
}
