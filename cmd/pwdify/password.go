package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type passwordKeyMap struct {
	Submit key.Binding
	Quit   key.Binding
}

func (k passwordKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Submit, k.Quit}
}

func (k passwordKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Submit, k.Quit},
	}
}

var passwordKeys = passwordKeyMap{
	Submit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "submit"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl-c", "quit"),
	),
}

type passwordModel struct {
	keys     passwordKeyMap
	password textinput.Model
	help     help.Model
	state    *state
}

func newPasswordModel(s *state) passwordModel {
	t := textinput.New()
	t.PromptStyle = alternateStyle
	t.TextStyle = secondaryStyle
	t.Cursor.Style = alternateStyle
	t.EchoMode = textinput.EchoPassword
	t.EchoCharacter = '*'
	t.CharLimit = 32
	t.Focus()

	return passwordModel{
		keys:     passwordKeys,
		password: t,
		help:     help.New(),
		state:    s,
	}
}

func (m passwordModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m passwordModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Submit):
			m.password.Blur()
			return m, func() tea.Msg {
				return PasswordCompleteMsg{Password: m.password.Value()}
			}
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}

	m.password, cmd = m.password.Update(msg)

	return m, cmd
}

func (m passwordModel) View() string {
	var b strings.Builder

	b.WriteString(primaryStyle.Margin(1, 2).Render("What password do you want to use?") + "\n")
	b.WriteString(lipgloss.NewStyle().Margin(0, 2).Render(m.password.View()) + "\n")
	b.WriteString(lipgloss.NewStyle().Margin(1, 2).Render(m.help.View(m.keys)) + "\n")

	return b.String()
}
