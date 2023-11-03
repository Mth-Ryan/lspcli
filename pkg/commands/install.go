package commands

import (
	"context"
	"time"

	"github.com/Mth-Ryan/lspcli/pkg/animations"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	d.animationWriter.Loading(ctx, "Loading...")

	return nil
}
