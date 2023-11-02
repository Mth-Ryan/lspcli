package commands

import (
	"strings"

	"github.com/Mth-Ryan/lspcli/internal/utils"
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/tools"
)

type ListCommand struct {
	reader tools.Reader
	writer tools.Writer
}

func NewListCommand(reader tools.Reader, writer tools.Writer) *ListCommand {
	return &ListCommand{
		reader,
		writer,
	}
}

// This can be replaced with a builder
func (l *ListCommand) createFilter(installed bool, kind string, lang string) func(models.Tool) bool {
	return func(t models.Tool) bool {
		pass := true

		if installed {
			pass = pass && t.LatestVersion != nil
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

	where := l.createFilter(installed, kind, lang)

	l.writer.Write(utils.Filter(tools, where))

	return nil
}
