package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Mth-Ryan/lspcli/config"
	"github.com/Mth-Ryan/lspcli/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func versionDialog() {
	conf := config.Get()

	fmt.Printf("lspcli %s", conf.Version)
	if conf.ExecMode == config.EXEC_DEBUG {
		fmt.Print(" \033[31mDEBUG Build\033[0m")
	}
	fmt.Print("\n")
}

func main() {
	version := flag.Bool("version", false, "Show the program version")
	flag.Parse()

	if *version {
		versionDialog()
		return
	}

	m := tui.InitTuiModel()
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
