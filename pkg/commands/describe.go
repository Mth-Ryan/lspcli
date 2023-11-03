package commands

import "github.com/Mth-Ryan/lspcli/pkg/tools"

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

	d.writer.Write(tool)

	return nil
}
