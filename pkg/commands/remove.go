package commands

import (
	"github.com/Mth-Ryan/lspcli/pkg/loggers"
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/providers"
	"github.com/Mth-Ryan/lspcli/pkg/result"
	"github.com/Mth-Ryan/lspcli/pkg/runtime"
	"github.com/Mth-Ryan/lspcli/pkg/tools"
)

type RemoveCommand struct {
	runtimeConf  *runtime.Conf
	reader       tools.Reader
	resultWriter result.Writer
	logger       loggers.Logger
}

func NewRemoveCommand(
	runtimeConf *runtime.Conf,
	reader tools.Reader,
	resultWriter result.Writer,
	logger loggers.Logger,
) *RemoveCommand {
	return &RemoveCommand{
		runtimeConf,
		reader,
		resultWriter,
		logger,
	}
}

func (d *RemoveCommand) Run(id string) error {
	tool, err := d.reader.Get(id)
	if err != nil {
		return err
	}

	provider, err := providers.GetProvider(tool, d.logger)
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
