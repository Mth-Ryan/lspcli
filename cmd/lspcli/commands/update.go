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
package commands

import (
	"github.com/Mth-Ryan/lspcli/cmd/lspcli/utils"
	"github.com/Mth-Ryan/lspcli/pkg/commands"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a installed tool",
	Long: `Update an installed tool to the last available version. Examples:
		
  lspcli update typescript-language-server
  lspcli update omnisharp`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.SupplyArgsOrHelp(cmd, args, 1)

		dependencies := utils.GetCommitActionDependencies(cmd)

		command := commands.NewUpdateCommand(
			dependencies.RuntimeConf,
			dependencies.ToolsReader,
			dependencies.ResultWriter,
			dependencies.Logger,
		)

		id := args[0]
		command.Run(id)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	updateCmd.Flags().BoolP("json", "j", false, "Send the output as a json object")
	updateCmd.Flags().BoolP("quiet", "q", false, "Supress the logging messages")
}
