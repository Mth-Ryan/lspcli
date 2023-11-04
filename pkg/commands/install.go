package commands

import (
	"context"
	"fmt"
	"time"

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go d.animationWriter.Loading(ctx, "installing")

	provider, err := providers.GetProvider(tool.Recipe.Kind)
	if err != nil {
		return err
	}

	var kind = models.RESULT_OK
	var message = fmt.Sprintf("%s was successfully", tool.Name)

	if err := provider.Install(tool); err != nil {
		kind = models.RESULT_ERR
		message = err.Error()
	}

	time.Sleep(3 * time.Second)

	d.resultWriter.Write(models.Result{
		Kind:    kind,
		Message: message,
	})

	return nil
}
