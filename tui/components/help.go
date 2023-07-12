package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var baseHelpStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#3f444a")).
	MarginLeft(1).
	MarginRight(1)

type KeyHelp struct {
	Keys string
	Description string
}

type HelpModel struct {
	keys []KeyHelp
}

func InitHelpModel(keys []KeyHelp) HelpModel {
	return HelpModel{ keys }
}

func (m HelpModel) View() string {
	viewList := make([]string, len(m.keys))
	for i, k := range m.keys {
		viewList[i] = k.Keys + " " + k.Description
	}

	return baseHelpStyle.Render(strings.Join(viewList, " • "))
}
