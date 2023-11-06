/*
Copyright © 2023 Mateus Ryan <mthryan@protonmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"

	"github.com/Mth-Ryan/lspcli/pkg/commands"
	"github.com/Mth-Ryan/lspcli/pkg/loggers"
	"github.com/Mth-Ryan/lspcli/pkg/result"
	"github.com/Mth-Ryan/lspcli/pkg/tools"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an installed tool",
	Long: `Remove an tool installed by the cli. Examples:

  lspcli remove typescript-language-server
  lspcli remove omnisharp		
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(1)
		}

		jsonOut, _ := cmd.Flags().GetBool("json")
		quietOut, _ := cmd.Flags().GetBool("quiet")
		runtimePath, _ := cmd.Flags().GetString("runtime")

		toolsReader := tools.NewRuntimeReader(runtimePath)
		var resultWriter result.Writer = result.NewPlainWriter()
		var logger loggers.Logger = loggers.NewStdOutLogger()

		if jsonOut {
			resultWriter = result.NewJsonWriter()
			logger = loggers.NewQuietLogger()
		} else if quietOut {
			logger = loggers.NewQuietLogger()
		}

		command := commands.NewRemoveCommand(toolsReader, resultWriter, logger)

		id := args[0]
		command.Run(id)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	removeCmd.Flags().BoolP("json", "j", false, "Send the output as a json object")
	removeCmd.Flags().BoolP("quiet", "q", false, "Supress the logging messages")
}
