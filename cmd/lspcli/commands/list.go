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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available tools",
	Long: `List all avaliable tools. Examples:
		
  lspcli list                   # list all tools
  lspcli list --lang typescript # list all tools for the typescript language
  lspcli list --kind lsp        # list all lsp tools
		
You also can also return the output as json with the json flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		dependencies := utils.GetShowActionDependencies(cmd)

		listCommand := commands.NewListCommand(
			*dependencies.RuntimeConf,
			dependencies.ToolsReader,
			dependencies.ToolsWriter,
		)

		installed, _ := cmd.Flags().GetBool("installed")
		kind, _ := cmd.Flags().GetString("kind")
		lang, _ := cmd.Flags().GetString("lang")

		err := listCommand.Run(installed, kind, lang)
		if err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().String("kind", "", "Show all tools for the specified kind")
	listCmd.Flags().String("lang", "", "Show all tools for the specified language")
	listCmd.Flags().Bool("installed", false, "Show only the installed tools")
	listCmd.Flags().BoolP("json", "j", false, "Send the output as a json object")
}
