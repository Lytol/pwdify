package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lytol/pwdify/pkg/pwdify"
)

var (
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

type NextMsg struct{}

func Next() tea.Msg {
	return NextMsg{}
}

type PasswordCompleteMsg struct {
	Password string
}

type FilesCompleteMsg struct {
	Files []string
}

func Run(cfg *pwdify.Config) error {
	_, err := tea.NewProgram(newModel(cfg)).Run()
	return err
}

type model struct {
	models  []tea.Model
	current int
	config  *pwdify.Config
}

func newModel(cfg *pwdify.Config) model {
	root := model{
		config: cfg,
		models: []tea.Model{
			newPasswordModel(),
			newFilesModel(cfg.Cwd),
			newStatusModel(cfg),
		},
		current: 0,
	}

	return root
}

func (m model) Current() tea.Model {
	return m.models[m.current]
}

func (m model) Init() tea.Cmd {
	return Next
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case NextMsg:
		return m.Next()
	case tea.WindowSizeMsg:
		cmds := make([]tea.Cmd, len(m.models))
		for i, mdl := range m.models {
			cm, cmd := mdl.Update(msg)
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
			m.models[i] = cm
		}
		return m, tea.Batch(cmds...)

	case PasswordCompleteMsg:
		m.config.Password = msg.Password
		return m, Next

	case FilesCompleteMsg:
		m.config.Files = msg.Files
		return m, Next
	}

	m.models[m.current], cmd = m.Current().Update(msg)

	return m, cmd
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString(m.Current().View())

	return b.String()
}

func (m model) Next() (model, tea.Cmd) {
	if m.config.Password == "" {
		m.current = 0
	} else if len(m.config.Files) == 0 {
		m.current = 1
	} else {
		m.current = 2
	}

	return m, m.Current().Init()
}
