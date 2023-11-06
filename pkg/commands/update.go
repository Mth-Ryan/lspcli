package commands

import (
	"github.com/Mth-Ryan/lspcli/pkg/loggers"
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/providers"
	"github.com/Mth-Ryan/lspcli/pkg/result"
	"github.com/Mth-Ryan/lspcli/pkg/tools"
)

type UpdateCommand struct {
	reader       tools.Reader
	resultWriter result.Writer
	logger       loggers.Logger
}

func NewUpdateCommand(reader tools.Reader, resultWriter result.Writer, logger loggers.Logger) *UpdateCommand {
	return &UpdateCommand{
		reader,
		resultWriter,
		logger,
	}
}

func (d *UpdateCommand) Run(id string) error {
	tool, err := d.reader.Get(id)
	if err != nil {
		return err
	}

	provider, err := providers.GetProvider(tool)
	if err != nil {
		return err
	}

	var kind = models.RESULT_OK
	var message = ""

	version, err := provider.LatestVersion()
	if err == nil {
		message = version
	} else {
		kind = models.RESULT_ERR
		message = err.Error()
	}

	d.resultWriter.Write(models.Result{
		Kind:    kind,
		Message: message,
	})

	return nil
}