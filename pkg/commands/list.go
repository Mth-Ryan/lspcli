package commands

import (
	"strings"

	"github.com/Mth-Ryan/lspcli/internal/utils"
	"github.com/Mth-Ryan/lspcli/pkg/handlers"
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/runtime"
	"github.com/Mth-Ryan/lspcli/pkg/tools"
)

type ListCommand struct {
	runtimeConf      runtime.Conf
	reader           tools.Reader
	writer           tools.Writer
	installedHandler handlers.InstallsListHandler
}

func NewListCommand(runtimeConf runtime.Conf, reader tools.Reader, writer tools.Writer) *ListCommand {
	installedHandler := handlers.NewJsonInstallsHandler(runtimeConf)

	return &ListCommand{
		runtimeConf,
		reader,
		writer,
		installedHandler,
	}
}

// This can be replaced with a builder
func (l *ListCommand) createFilter(installed bool, kind string, lang string) func(models.Tool) bool {
	return func(t models.Tool) bool {
		pass := true

		if installed {
			pass = pass && t.InstalledVersion != nil
		}

		if kind != "" {
			pass = pass && strings.ToLower(t.Kind) == strings.ToLower(kind)
		}

		if lang != "" {
			pass = pass && utils.Any(
				utils.Map(t.Languages, strings.ToLower),
				strings.ToLower(lang),
			)
		}

		return pass
	}
}

func (l *ListCommand) Run(installed bool, kind string, lang string) error {
	tools, err := l.reader.GetAll()
	if err != nil {
		return err
	}

	installedList, err := l.installedHandler.GetInstalls()
	if err != nil {
		return err
	}

	where := l.createFilter(installed, kind, lang)

	newTools := utils.Map(tools, func(t models.Tool) models.Tool {
		newTool := t
		if version, ok := installedList[t.ID]; ok {
			newTool.InstalledVersion = version
		} else {
			newTool.InstalledVersion = nil
		}

		return newTool
	})

	l.writer.WriteAll(utils.Filter(newTools, where))

	return nil
}
