package tui

import (
	"fmt"
	"os"
	"strings"

	"github.com/allandlobr/go-bookmark-cli/utils"
	"github.com/atotto/clipboard"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const DB_FILENAME string = "db.json"

type Model struct {
	items         []string // items on the list
	cursor        int      // which item list item our cursor is pointing at
	title         string   // title
	showingGroups bool     // the items being shown are groups
}

func initModel(model Model) Model {
	return model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// What was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c":
			return m, tea.Quit

		// Exit the program and save any change in bookmarks.
		case "q":
			if !m.showingGroups {
				updateGroup(strings.ToLower(m.title), m.items)
			}
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}

		// Deletes bookmark from list. Only saves if exited with q key.
		case "d":
			if !m.showingGroups && (len(m.items) > 0) {
				m.items = append(m.items[:m.cursor], m.items[m.cursor+1:]...)
			}

		// Selects item on the list. If is a group, it shows the bookmarks inside it, if it is a bookmark, a copy is made to clipboard.
		case "enter":
			if m.showingGroups && (len(m.items) > 0) {
				groupName := strings.ToLower(m.items[m.cursor])
				groups := utils.ReadJSON(DB_FILENAME)
				m.items = groups[groupName]
				m.title = strings.ToUpper(groupName)
				m.showingGroups = false
				break
			}

			if !m.showingGroups && (len(m.items) > 0) {
				clipboard.WriteAll(m.items[m.cursor])
				return m, tea.Quit
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
func (m Model) View() string {
	var (
		header, body, footer, fullList string
	)
	// The header
	header = lipgloss.NewStyle().AlignHorizontal(lipgloss.Center).Inline(true).Width(40).Bold(true).SetString(m.title).String()

	// The body

	// Iterate over our choices
	for i, choice := range m.items {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
			body = lipgloss.JoinVertical(lipgloss.Left, body, lipgloss.NewStyle().Background(lipgloss.Color("#3b16b5")).Padding(0, 1).SetString(fmt.Sprintf("%s %d. %s", cursor, i+1, choice)).String())
		} else {
			body = lipgloss.JoinVertical(lipgloss.Left, body, lipgloss.NewStyle().SetString(fmt.Sprintf("%s %d. %s", cursor, i+1, choice)).String())
		}
	}

	// The footer
	footerStyle := lipgloss.NewStyle().AlignHorizontal(lipgloss.Left).Padding(0, 1).Width(40).Foreground(lipgloss.Color("#3C3C3C"))
	if m.showingGroups {
		footer = lipgloss.JoinVertical(lipgloss.Left, footer,
			footerStyle.SetString("Press q to quit.").String(),
			footerStyle.SetString("Press enter to list group.").String())
	} else {
		footer = lipgloss.JoinVertical(lipgloss.Left, footer,
			footerStyle.SetString("Press q to quit.").String(),
			footerStyle.SetString("Press d to delete.").String(),
			footerStyle.SetString("Press enter to copy.").String())
	}

	fullList = lipgloss.JoinVertical(lipgloss.Left, header, body, footer)

	// Send the UI for rendering
	return lipgloss.NewStyle().Margin(0, 4, 1).Border(lipgloss.NormalBorder()).Render(fullList)
}

func updateGroup(groupName string, items []string) error {

	groups := utils.ReadJSON(DB_FILENAME)
	groups[groupName] = items

	err := utils.WriteJSON(DB_FILENAME, groups)
	if err != nil {
		return err
	}

	return nil

}

func StartTui(items []string, title string, showingGroups bool) {
	initialProps := Model{
		items:         items,
		title:         title,
		cursor:        0,
		showingGroups: showingGroups,
	}

	if _, err := tea.NewProgram(initModel(initialProps)).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
