package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/lytol/pwdify/internal/util"
)

var (
	logger util.Logger

	primaryStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("#B9EBFF"))
	secondaryStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("#65C1E3"))
	helpStyle               = lipgloss.NewStyle().Foreground(lipgloss.Color("#777777"))
	alternateStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("#DB9655"))
	alternateSecondaryStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#BF9D80"))
)

type model struct {
	password textinput.Model
}

func initialModel() model {
	t := textinput.New()
	t.PromptStyle = alternateStyle
	t.TextStyle = secondaryStyle
	t.Cursor.Style = alternateStyle
	t.Placeholder = "Password"
	t.EchoMode = textinput.EchoPassword
	t.EchoCharacter = '*'
	t.CharLimit = 32
	t.Focus()

	return model{
		password: t,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		logger.Logf("Update | key: `%s`\n", msg.String())

		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEnter:
			logger.Logf("Update | quit\n")
			m.password.Blur()
			return m, tea.Quit
		}
	}

	m.password, cmd = m.password.Update(msg)

	return m, cmd
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString(primaryStyle.MarginBottom(1).Render("What password do you want to use?") + "\n")
	b.WriteString(m.password.View() + "\n")
	b.WriteString(helpStyle.MarginTop(1).Render("enter to submit, ctrl-c to quit") + "\n")

	return b.String()
}

func main() {
	var err error

	logger, err = util.NewLogger()
	if err != nil {
		fmt.Printf("could not start logger: %s\n", err)
		os.Exit(1)
	}

	if _, err = tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
