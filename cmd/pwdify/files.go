package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

type filesModel struct {
	filepicker filepicker.Model
}

func newFilesModel() filesModel {
	fp := filepicker.New()
	fp.AllowedTypes = []string{".html"}
	fp.FileAllowed = true
	fp.DirAllowed = false
	fp.ShowPermissions = false
	fp.ShowSize = false
	fp.Height = 5

	fp.Styles.Cursor = alternateStyle
	fp.Styles.Selected = secondaryStyle
	fp.Styles.Directory = tertiaryStyle

	return filesModel{
		filepicker: fp,
	}
}

func (m filesModel) Init() tea.Cmd {
	logger.Logf("Init[files]\n")
	return m.filepicker.Init()
}

func (m filesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.filepicker, cmd = m.filepicker.Update(msg)

	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
		logger.Logf("Update[files] | selected file: `%s`\n", path)
		return m, func() tea.Msg {
			return FilesCompleteMsg{Path: path}
		}
	}

	return m, cmd
}

func (m filesModel) View() string {
	var b strings.Builder

	b.WriteString(primaryStyle.MarginBottom(1).Render("What file do you want to password protect?") + "\n")
	b.WriteString(m.filepicker.View() + "\n")

	return b.String()
}
