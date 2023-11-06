package utils

import (
	"os"

	"github.com/Mth-Ryan/lspcli/pkg/loggers"
	"github.com/Mth-Ryan/lspcli/pkg/result"
	"github.com/Mth-Ryan/lspcli/pkg/runtime"
	"github.com/Mth-Ryan/lspcli/pkg/tools"
	"github.com/spf13/cobra"
)

func SupplyArgsOrHelp(cmd *cobra.Command, args []string, required int) {
	if len(args) < required {
		cmd.Help()
		os.Exit(1)
	}
}

type CommitActionDependencies struct {
	RuntimeConf  *runtime.Conf
	ToolsReader  tools.Reader
	ResultWriter result.Writer
	Logger       loggers.Logger
}

func GetCommitActionDependencies(cmd *cobra.Command) CommitActionDependencies {
	quietOut, _ := cmd.Flags().GetBool("quiet")
	jsonOut, _ := cmd.Flags().GetBool("json")
	runtimePath, _ := cmd.Flags().GetString("runtime")

	dependencies := CommitActionDependencies{
		RuntimeConf:  runtime.NewConf(runtimePath),
		ToolsReader:  tools.NewRuntimeReader(runtimePath),
		ResultWriter: result.NewPlainWriter(),
		Logger:       loggers.NewStdOutLogger(),
	}

	if jsonOut {
		dependencies.ResultWriter = result.NewJsonWriter()
		dependencies.Logger = loggers.NewQuietLogger()
	} else if quietOut {
		dependencies.Logger = loggers.NewQuietLogger()
	}

	return dependencies
}

type ShowActionDependencies struct {
	RuntimeConf *runtime.Conf
	ToolsReader tools.Reader
	ToolsWriter tools.Writer
}

func GetShowActionDependencies(cmd *cobra.Command) ShowActionDependencies {
	jsonOut, _ := cmd.Flags().GetBool("json")
	runtimePath, _ := cmd.Flags().GetString("runtime")

	dependencies := ShowActionDependencies{
		RuntimeConf: runtime.NewConf(runtimePath),
		ToolsReader: tools.NewRuntimeReader(runtimePath),
		ToolsWriter: tools.NewTableWriter(),
	}

	if jsonOut {
		dependencies.ToolsWriter = tools.NewJsonWriter()
	}

	return dependencies
}
