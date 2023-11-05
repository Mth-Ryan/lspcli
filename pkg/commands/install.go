package commands

import (
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/providers"
	"github.com/Mth-Ryan/lspcli/pkg/result"
	"github.com/Mth-Ryan/lspcli/pkg/tools"
)

type InstallCommand struct {
	reader       tools.Reader
	resultWriter result.Writer
}

func NewInstallCommand(reader tools.Reader, resultWriter result.Writer) *InstallCommand {
	return &InstallCommand{
		reader,
		resultWriter,
	}
}

func (d *InstallCommand) Run(id string) error {
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
