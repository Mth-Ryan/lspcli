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
	"github.com/Mth-Ryan/lspcli/cmd/utils"
	"github.com/Mth-Ryan/lspcli/pkg/commands"
	"github.com/spf13/cobra"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe a tool",
	Long: `Describe a tool. You can check all informations
about the tool; latest version, installed version, dependencies,
description, etc. Examples:

  lspcli describe typescript-language-server
  lspcli describe omnisharp		
`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.SupplyArgsOrHelp(cmd, args, 1)

		dependencies := utils.GetShowActionDependencies(cmd)

		describeCommand := commands.NewDescribeCommand(
			dependencies.ToolsReader,
			dependencies.ToolsWriter,
		)

		id := args[0]
		err := describeCommand.Run(id)
		if err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// describeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// describeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	describeCmd.Flags().BoolP("json", "j", false, "Send the output as a json object")
}
