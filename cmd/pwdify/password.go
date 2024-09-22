package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type passwordModel struct {
	password textinput.Model
	state    *state
}

func newPasswordModel(s *state) passwordModel {
	t := textinput.New()
	t.PromptStyle = alternateStyle
	t.TextStyle = secondaryStyle
	t.Cursor.Style = alternateStyle
	t.Placeholder = "Password"
	t.EchoMode = textinput.EchoPassword
	t.EchoCharacter = '*'
	t.CharLimit = 32
	t.Focus()

	return passwordModel{
		state:    s,
		password: t,
	}
}

func (m passwordModel) Init() tea.Cmd {
	logger.Logf("Init[password]\n")
	return textinput.Blink
}

func (m passwordModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		logger.Logf("Update[password] | key: `%s`\n", msg.String())

		switch msg.Type {
		case tea.KeyEnter:
			logger.Logf("Update[password] | enter\n")
			m.password.Blur()
			return m, func() tea.Msg {
				return PasswordCompleteMsg{Password: m.password.Value()}
			}
		}
	}

	m.password, cmd = m.password.Update(msg)

	return m, cmd
}

func (m passwordModel) View() string {
	var b strings.Builder

	b.WriteString(primaryStyle.MarginLeft(2).MarginBottom(1).Render("What password do you want to use?") + "\n")
	b.WriteString(lipgloss.NewStyle().MarginLeft(2).Render(m.password.View()) + "\n")

	return b.String()
}
