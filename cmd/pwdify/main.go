package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/lytol/pwdify"
	"github.com/lytol/pwdify/internal/util"
)

var (
	logger util.Logger

	primaryStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("#B9EBFF"))
	secondaryStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("#65C1E3"))
	tertiaryStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("#208EAD"))
	helpStyle               = lipgloss.NewStyle().Foreground(lipgloss.Color("#777777"))
	alternateStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("#DB9655"))
	alternateSecondaryStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#BF9D80"))
)

type PasswordCompleteMsg struct {
	Password string
}

type FilesCompleteMsg struct {
	Path string
}

type model struct {
	models  []tea.Model
	current int
	config  pwdify.Config
}

func newModel() model {
	return model{
		models: []tea.Model{
			newPasswordModel(),
			newFilesModel(),
		},
		current: 0,
		config:  pwdify.Config{},
	}
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
	case tea.KeyMsg:
		logger.Logf("Update | key: `%s`\n", msg.String())

		switch msg.Type {
		case tea.KeyCtrlC:
			logger.Logf("Update | quit\n")
			return m, tea.Quit
		}

	case PasswordCompleteMsg:
		logger.Logf("Update | password: `%s`\n", msg.Password)
		m.config.Password = msg.Password
		m.current += 1
		return m, m.Current().Init()

	case FilesCompleteMsg:
		logger.Logf("Update | file: `%s`\n", msg.Path)
		m.config.Path = msg.Path
		return m, tea.Quit
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

	if _, err = tea.NewProgram(newModel()).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
