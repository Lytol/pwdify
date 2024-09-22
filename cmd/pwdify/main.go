package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/lytol/pwdify/internal/util"
)

var (
	logger util.Logger

	primaryColor            = "#B9EBFF"
	secondaryColor          = "#65C1E3"
	tertiaryColor           = "#208EAD"
	alternateColor          = "#DB9655"
	alternateSecondaryColor = "#BF9D80"

	primaryStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color(primaryColor))
	secondaryStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color(secondaryColor))
	tertiaryStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color(tertiaryColor))
	alternateStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color(alternateColor))
	alternateSecondaryStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(alternateSecondaryColor))

	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#5FD35F"))
	failureStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#D75F5F"))
)

type PasswordCompleteMsg struct {
	Password string
}

type FilesCompleteMsg struct {
	Files []string
}

type model struct {
	models  []tea.Model
	current int
	state   *state
}

func newModel(s *state) model {
	root := model{
		state: s,
		models: []tea.Model{
			newPasswordModel(s),
			newFilesModel(s),
			newStatusModel(s),
		},
		current: 0,
	}

	return root
}

func (m model) Current() tea.Model {
	return m.models[m.current]
}

func (m model) Init() tea.Cmd {
	logger.Logf("Init\n")
	return m.Current().Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		for _, m := range m.models {
			m.Update(msg)
		}
		return m, nil

	case tea.KeyMsg:
		logger.Logf("Update | key: `%s`\n", msg.String())

		switch msg.Type {
		case tea.KeyCtrlC:
			logger.Logf("Update | quit\n")
			return m, tea.Quit
		}

	case PasswordCompleteMsg:
		logger.Logf("Update | password: `%s`\n", msg.Password)
		m.state.password = msg.Password
		m.current += 1
		return m, m.Current().Init()

	case FilesCompleteMsg:
		logger.Logf("Update | file: `%s`\n", msg.Files)
		m.state.files = msg.Files
		m.current += 1
		return m, m.Current().Init()
	}

	m.models[m.current], cmd = m.Current().Update(msg)

	return m, cmd
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString(m.Current().View())

	return b.String()
}

func main() {
	var err error

	logger, err = util.NewLogger()
	if err != nil {
		fmt.Printf("could not start logger: %s\n", err)
		os.Exit(1)
	}

	state := &state{
		password: "",
		files:    []string{},
	}

	if _, err = tea.NewProgram(newModel(state)).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
