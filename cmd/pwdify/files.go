package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type fileItem struct {
	path     string
	trimmed  string
	selected bool
}

func (fi fileItem) FilterValue() string {
	return fi.trimmed
}

type fileItemDelegate struct{}

func (d fileItemDelegate) Height() int                             { return 1 }
func (d fileItemDelegate) Spacing() int                            { return 0 }
func (d fileItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d fileItemDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	i, ok := item.(*fileItem)
	if !ok {
		return
	}

	checkbox := "☐"

	if i.selected {
		checkbox = "■"
	}

	if index == m.Index() {
		fmt.Fprint(w, strings.Join([]string{
			alternateStyle.MarginLeft(2).Render("•"),
			alternateSecondaryStyle.Render(checkbox),
			secondaryStyle.Render(i.trimmed),
		}, " "))
	} else {
		fmt.Fprint(w, strings.Join([]string{
			alternateSecondaryStyle.MarginLeft(4).Render(checkbox),
			tertiaryStyle.Render(i.trimmed),
		}, " "))
	}
}

type readDirMsg struct {
	Files []string
}

type filesModel struct {
	files list.Model
	state *state
}

func newFilesModel(s *state) filesModel {
	l := list.New([]list.Item{}, fileItemDelegate{}, 80, 10)
	l.Title = "Select files to password protect"
	l.Styles.Title = primaryStyle
	l.SetHeight(12)

	return filesModel{
		state: s,
		files: l,
	}
}

func (m filesModel) Init() tea.Cmd {
	return readDir(m.state.cwd)
}

func (m filesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case readDirMsg:
		// Clear existing items
		for i := range m.files.Items() {
			m.files.RemoveItem(i)
		}

		// Add new items
		for i, f := range msg.Files {
			trimmed := strings.TrimPrefix(f, m.state.cwd+string(filepath.Separator))

			m.files.InsertItem(i, &fileItem{
				path:    f,
				trimmed: trimmed,
			})
		}

		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeySpace:
			i, ok := m.files.SelectedItem().(*fileItem)
			if ok {
				i.selected = !i.selected
			}
			return m, nil
		case tea.KeyEnter:
			selectedFiles := []string{}

			for _, item := range m.files.Items() {
				i, ok := item.(*fileItem)
				if !ok {
					return m, tea.Quit
				}

				if i.selected {
					selectedFiles = append(selectedFiles, i.path)
				}
			}

			return m, func() tea.Msg {
				return FilesCompleteMsg{Files: selectedFiles}
			}
		}
	}

	m.files, cmd = m.files.Update(msg)
	return m, cmd
}

func (m filesModel) View() string {
	var b strings.Builder

	b.WriteString(lipgloss.NewStyle().Margin(1, 2).Render(m.files.View()))

	return b.String()
}

func readDir(path string) tea.Cmd {
	return func() tea.Msg {
		files := []string{}

		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".html" {
				files = append(files, path)
			}
			return err
		})

		if err != nil {
			fmt.Fprintf(os.Stderr, "could not read directory: %s\n", err)
			return tea.Quit
		}

		return readDirMsg{Files: files}
	}
}
