package tui

import (
	"github.com/Mth-Ryan/lspcli/config"
	"github.com/Mth-Ryan/lspcli/core"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	_TABLE int = iota
)

type model struct {
	state int

	tablePage tablePageModel
}

func InitTuiModel() model {
	recipes, err := core.LoadAndParseRecipes(config.Get().RecipesFilePath)
	if err != nil {
		panic(err)
	}

	return model{
		state:  _TABLE,

		tablePage: InitTablePageModel(recipes),
	}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var tcmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	m.tablePage, tcmd = m.tablePage.Update(msg)
	return m, tea.Batch(tcmd)
}

func (m model) View() string {
	switch m.state {
	case _TABLE:
		return m.tablePage.View()

	default:
		return ""
	}
}
