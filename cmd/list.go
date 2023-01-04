/*
Copyright © 2022 allandlobr allandlobr@gmail.com
*/
package cmd

import (
	"strings"

	"github.com/allandlobr/go-bookmark-cli/tui"
	"github.com/allandlobr/go-bookmark-cli/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists bookmarks.",
	Long:  `Lists bookmarks.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			group []string
			items []string
		)
		groups := utils.ReadJSON(DB_FILENAME)

		if len(args) == 0 {

			for k := range groups {
				items = append(items, k)
			}

			tui.StartTui(items, "GROUP LIST", true)

		} else {
			groupName := strings.ToLower(args[0])

			group = groups[groupName]

			items = append(items, group...)

			tui.StartTui(items, strings.ToUpper(groupName), false)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
