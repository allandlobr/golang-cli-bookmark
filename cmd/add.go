/*
Copyright © 2022 allandlobr allandlobr@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/allandlobr/go-bookmark-cli/utils"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(0), cobra.MaximumNArgs(2)),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			group string
			link  string
		)

		switch len(args) {
		case 0:
			group = "main"
			link, _ = clipboard.ReadAll()
		case 1:
			group = args[0]
			link, _ = clipboard.ReadAll()
		case 2:
			group = args[0]
			link = args[1]
		}

		groups := utils.ReadJSON(DB_FILENAME)

		for k := range groups {
			if k == group {
				groups[k] = append(groups[k], link)

				err := utils.WriteJSON(DB_FILENAME, groups)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				fmt.Println("Bookmark successfull!")
				os.Exit(0)
			}
		}

		fmt.Println("Group specified does not exist!")
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
