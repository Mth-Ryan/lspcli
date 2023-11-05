package commands

import (
	"github.com/Mth-Ryan/lspcli/pkg/providers"
	"github.com/Mth-Ryan/lspcli/pkg/tools"
)

type DescribeCommand struct {
	reader tools.Reader
	writer tools.Writer
}

func NewDescribeCommand(reader tools.Reader, writer tools.Writer) *DescribeCommand {
	return &DescribeCommand{
		reader,
		writer,
	}
}

func (d *DescribeCommand) Run(id string) error {
	tool, err := d.reader.Get(id)
	if err != nil {
		return err
	}

	provider, err := providers.GetProvider(tool)
	if err != nil {
		return err
	}

	latestVersion, err := provider.LatestVersion()
	if err != nil {
		return err
	}

	installedVersion, err := provider.InstaledVersion()
	if err != nil {
		return err
	}

	tool.LatestVersion = &latestVersion
	if installedVersion != "" {
		tool.InstalledVersion = &installedVersion
	}

	d.writer.Write(tool)

	return nil
}
