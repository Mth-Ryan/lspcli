package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/Mth-Ryan/lspcli/tui"
)

const VERSION = "0.1.0"
const RECIPES = "runtime/recipes.yml"

func versionDialog() {
	fmt.Printf("lspcli %s\n", VERSION)
}

func main() {
	m := tui.InitTuiModel()
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
