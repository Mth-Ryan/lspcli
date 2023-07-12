package tui

import (
	"github.com/Mth-Ryan/lspcli/core"
	"github.com/Mth-Ryan/lspcli/tui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type tablePageModel struct {
    lspTable components.LspTableModel
	bar components.BarModel
	help components.HelpModel
}

func InitTablePageModel(recipes []core.Recipe) tablePageModel {
	return tablePageModel{
		lspTable: components.InitTable(recipes),
		bar: components.InitBarModel(),
		help: components.InitHelpModel([]components.KeyHelp{
			{"k/↑", "move up"},
			{"j/↓", "move down"},
			{"↲", "open lsp options"},
		}),
	}
}

func (m tablePageModel) Init() tea.Cmd { return nil }

func (m tablePageModel) Update(msg tea.Msg) (tablePageModel, tea.Cmd) {
	var tcmd tea.Cmd
	var bcmd tea.Cmd

	m.lspTable, tcmd = m.lspTable.Update(msg)
	m.bar, bcmd = m.bar.Update(msg)

	return m, tea.Batch(tcmd, bcmd)
}

func (m tablePageModel) View() string {
	return m.lspTable.View() + "\n" + m.bar.View() + "\n" + m.help.View() + "\n"
}
