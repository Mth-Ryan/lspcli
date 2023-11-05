package commands

import (
	"context"

	"github.com/Mth-Ryan/lspcli/pkg/animations"
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/providers"
	"github.com/Mth-Ryan/lspcli/pkg/result"
	"github.com/Mth-Ryan/lspcli/pkg/tools"
)

type InstallCommand struct {
	reader          tools.Reader
	resultWriter    result.Writer
	animationWriter animations.Writer
}

func NewInstallCommand(reader tools.Reader, resultWriter result.Writer, animationWriter animations.Writer) *InstallCommand {
	return &InstallCommand{
		reader,
		resultWriter,
		animationWriter,
	}
}

func (d *InstallCommand) Run(id string) error {
	tool, err := d.reader.Get(id)
	if err != nil {
		return err
	}

	ctx, animationCancel := context.WithCancel(context.Background())
	defer animationCancel()
	defer d.animationWriter.Clear()

	go d.animationWriter.Loading(ctx, "installing")

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

	animationCancel()
	d.animationWriter.Clear()

	d.resultWriter.Write(models.Result{
		Kind:    kind,
		Message: message,
	})

	return nil
}
