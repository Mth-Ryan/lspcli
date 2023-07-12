package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var firstBlockStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#5689F0")).
	Foreground(lipgloss.Color("#DFDFDF")).
	PaddingLeft(1).
	PaddingRight(1).
	MarginLeft(1).
	Bold(true)

var middleStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#23272e")).
	Foreground(lipgloss.Color("#DFDFDF")).
	Width(12)

var rightBlockStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#151619")).
	Foreground(lipgloss.Color("#78BD65")).
	PaddingLeft(1).
	PaddingRight(1)

var lastBlockStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#EE7B29")).
	Foreground(lipgloss.Color("#151619")).
	MarginRight(1).
	PaddingLeft(1).
	PaddingRight(1)

type BarModel struct {}

func InitBarModel() BarModel {
	return BarModel{}
}

func (m BarModel) Init() tea.Cmd {
	return nil
}

func (m BarModel) Update(msg tea.Msg) (BarModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		middleStyle.Width(msg.Width - (2 + 9 + 3 + 7))
	}

	return m, cmd
}

func (m BarModel) View() string {
	return firstBlockStyle.Render("Lsp-cli") +
		middleStyle.Render("") +
		rightBlockStyle.Render("✓") +
		lastBlockStyle.Render("UTF-8")
}
