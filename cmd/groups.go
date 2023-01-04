/*
Copyright © 2022 allandlobr allandlobr@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/allandlobr/go-bookmark-cli/utils"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// groupsCmd represents the groups command
var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var groupsAddCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(1),
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			groupName = args[0]
		)

		groups := utils.ReadJSON(DB_FILENAME)

		for k := range groups {
			if k == groupName {
				fmt.Println("Group name already exists!")
				os.Exit(1)
			}
		}

		groups[groupName] = []string{}

		err := utils.WriteJSON(DB_FILENAME, groups)

		if err != nil {
			os.Exit(1)
		}

		fmt.Println("Group added!")
		os.Exit(0)
	},
}

var groupsRemoveCmd = &cobra.Command{
	Use:   "remove",
	Args:  cobra.ExactArgs(1),
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			groupName = args[0]
		)

		groups := utils.ReadJSON(DB_FILENAME)

		for k := range groups {
			if k == groupName {
				delete(groups, groupName)
				err := utils.WriteJSON(DB_FILENAME, groups)
				if err != nil {
					os.Exit(1)
				}
				os.Exit(0)
			}
		}

		fmt.Println("Group doesn't exist!")
		os.Exit(1)
	},
}

var groupsListCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		groups := utils.ReadJSON(DB_FILENAME)

		if len(groups) == 0 {
			os.Exit(1)
		}

		title := lipgloss.NewStyle().Bold(true).AlignHorizontal(lipgloss.Center).MarginBottom(1).Width(40).Render("GROUPS")
		var b string
		theader := lipgloss.JoinHorizontal(lipgloss.Left, lipgloss.NewStyle().Width(20).Border(lipgloss.NormalBorder(), false, false, true, false).AlignHorizontal(lipgloss.Center).Bold(true).Render("NAME"), lipgloss.NewStyle().Width(20).Border(lipgloss.NormalBorder(), false, false, true, false).AlignHorizontal(lipgloss.Center).Bold(true).Render("ITEMS"))

		b = theader
		for k, v := range groups {
			row := lipgloss.JoinHorizontal(lipgloss.Left, lipgloss.NewStyle().Width(20).AlignHorizontal(lipgloss.Center).Bold(true).Render(""+k), lipgloss.NewStyle().Width(20).AlignHorizontal(lipgloss.Center).Bold(true).Render(fmt.Sprintf("%d", len(v))))
			b = lipgloss.JoinVertical(lipgloss.Left, b, row)
		}

		result := lipgloss.NewStyle().Margin(0, 4, 1).Border(lipgloss.NormalBorder()).Render(lipgloss.JoinVertical(lipgloss.Left, title, b, ""))
		fmt.Println(result)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(groupsCmd)
	groupsCmd.AddCommand(groupsAddCmd)
	groupsCmd.AddCommand(groupsRemoveCmd)
	groupsCmd.AddCommand(groupsListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groupsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
