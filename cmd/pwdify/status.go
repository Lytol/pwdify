package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lytol/pwdify/pkg/pwdify"
)

type statusTickMsg struct{}

type statusModel struct {
	progress progress.Model
	spinner  spinner.Model
	state    *state
	engine   *pwdify.Engine
}

func newStatusModel(s *state) statusModel {
	prg := progress.New(progress.WithSolidFill(tertiaryColor))

	spn := spinner.New()
	spn.Spinner = spinner.Dot
	spn.Style = alternateStyle.Width(4)

	return statusModel{
		progress: prg,
		spinner:  spn,
		engine:   pwdify.New(),
		state:    s,
	}
}

func (m statusModel) Init() tea.Cmd {
	return tea.Batch(m.encrypt, m.spinner.Tick)
}

func (m statusModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		logger.Logf("Update[status] | key: `%s`\n", msg.String())
		return m, nil

	case statusTickMsg:
		logger.Logf("Update[status] | tick\n")
		tick := m.tick()
		progress := m.progress.SetPercent(m.state.PercentCompleted())
		return m, tea.Batch(tick, progress)

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

	statusStr := fmt.Sprintf("%d/%d completed • %d errors\n", m.state.CompleteCount(), m.state.TotalCount(), m.state.ErrorCount())

	if !m.state.Completed() {
		b.WriteString(lipgloss.NewStyle().MarginLeft(2).Width(4).Render(m.spinner.View()))
	} else if m.state.ErrorCount() > 0 {
		b.WriteString(failureStyle.MarginLeft(2).Width(4).Render("✗"))
	} else {
		b.WriteString(successStyle.MarginLeft(2).Width(4).Render("✔"))
	}

	b.WriteString(primaryStyle.Render(statusStr) + "\n")

	return b.String()
}

func (m statusModel) encrypt() tea.Msg {
	m.state.ch = m.engine.Run(m.state.files, m.state.password)
	return statusTickMsg{}
}

func (m statusModel) tick() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		select {
		case s, ok := <-m.state.ch:
			if !ok {
				return tea.QuitMsg{}
			}

			m.state.status = append(m.state.status, s)
			if s.Error != nil {
				logger.Logf("encrypt[%s] | error: %s\n", s.File, s.Error)
			} else {
				logger.Logf("encrypt[%s] | done\n", s.File)
			}

			// TODO: update progress
			// TODO: update list
			return statusTickMsg{}
		default:
			return statusTickMsg{}
		}
	})
}
