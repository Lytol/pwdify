package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lytol/pwdify/pkg/pwdify"
)

type startMsg struct{}

func start() tea.Msg {
	return startMsg{}
}

type statusModel struct {
	progress progress.Model
	spinner  spinner.Model
	state    *state
	engine   *pwdify.Engine

	status chan pwdify.Status

	total     int
	completed int
	errors    int
}

func newStatusModel(s *state) statusModel {
	prg := progress.New(progress.WithSolidFill(tertiaryColor))

	spn := spinner.New()
	spn.Spinner = spinner.Dot
	spn.Style = alternateStyle.Width(4)

	return statusModel{
		progress: prg,
		spinner:  spn,
		state:    s,
	}
}

func (m statusModel) Init() tea.Cmd {
	return tea.Batch(start, m.spinner.Tick)
}

func (m statusModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case startMsg:
		var err error
		m.engine, err = pwdify.New(m.state.files, m.state.password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not start pwdify engine: %s\n", err)
			return m, tea.Quit
		}
		m.total = len(m.engine.Files)
		m.status = m.engine.Run()
		logger.Logf("status[files] | %+v\n", m.engine.Files)
		return m, m.tick()

	case pwdify.Status:
		m.completed += 1

		if msg.Error != nil {
			m.errors += 1
		}

		progress := m.progress.SetPercent(m.percentComplete())
		return m, tea.Batch(m.tick(), progress)

	case progress.FrameMsg:
		pm, cmd := m.progress.Update(msg)
		m.progress = pm.(progress.Model)
		return m, cmd

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m statusModel) View() string {
	var b strings.Builder

	b.WriteString(lipgloss.NewStyle().Margin(1, 2).Render(m.progress.View()) + "\n")

	statusStr := fmt.Sprintf("%d/%d completed • %d errors\n", m.completed, m.total, m.errors)

	if !m.finished() {
		b.WriteString(lipgloss.NewStyle().MarginLeft(2).Width(4).Render(m.spinner.View()))
	} else if m.errors > 0 {
		b.WriteString(failureStyle.MarginLeft(2).Width(4).Render("✗"))
	} else {
		b.WriteString(successStyle.MarginLeft(2).Width(4).Render("✔"))
	}

	b.WriteString(primaryStyle.Render(statusStr) + "\n")

	return b.String()
}

func (m statusModel) tick() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		s, ok := <-m.status
		if !ok {
			return tea.QuitMsg{}
		}
		return s
	})
}

func (m statusModel) percentComplete() float64 {
	return float64(m.completed) / float64(m.total)
}

func (m statusModel) finished() bool {
	return m.completed == m.total
}
