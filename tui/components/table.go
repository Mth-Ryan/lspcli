package components

import (
	"fmt"

	"github.com/Mth-Ryan/lspcli/core"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var tableStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#3f444a")).
	MarginLeft(1).
	MarginRight(1)

type LspTableModel struct {
	table table.Model
}

func getColumnsNormalized(totalCols int) []table.Column {
	lspNameWidth := totalCols - (3 + 20 + 9 + 6 + 6) // total - (Id + Lang + Installed + Borders + gaps)
	if lspNameWidth < 12 {
		lspNameWidth = 12
	}

	return []table.Column{
		{Title: "Id", Width: 3},
		{Title: "Lsp", Width: lspNameWidth},
		{Title: "Language ↓", Width: 20},
		{Title: "Installed", Width: 9},
	}
}

func InitTable(recipes []core.Recipe) LspTableModel {
	columns := getColumnsNormalized(70)

	rows := []table.Row{}
	for idx, recipe := range recipes {
		id := fmt.Sprintf("%03d", idx + 1)
		rows = append(rows, []string{id, recipe.Lsp, recipe.Lang, "      " + "[ ]"})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		Foreground(lipgloss.Color("#3f444a")).
		BorderForeground(lipgloss.Color("#3f444a")).
		BorderBottom(true).
		Bold(true)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#DFDFDF")).
		Background(lipgloss.Color("#5689F0")).
		Bold(true)

	t.SetStyles(s)

	return LspTableModel{
		table: t,
	}
}

func (m LspTableModel) Init() tea.Cmd { return nil }

func (m LspTableModel) Update(msg tea.Msg) (LspTableModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.table.SetColumns(getColumnsNormalized(msg.Width))
		m.table.SetHeight(msg.Height - 8)
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m LspTableModel) View() string {
	return tableStyle.Render(m.table.View())
}
