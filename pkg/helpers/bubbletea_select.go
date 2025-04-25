package helpers

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Item represents a selectable item
type item struct {
	title string
	desc  string
}

// Title implements the list.Item interface
func (i item) Title() string { return i.title }

// Description implements the list.Item interface
func (i item) Description() string { return i.desc }

// FilterValue implements the list.Item interface
func (i item) FilterValue() string { return i.title }

// Model represents the Bubbletea model for selection
type selectionModel struct {
	list     list.Model
	choice   string
	quitting bool
}

// Init implements the Bubbletea model interface
func (m selectionModel) Init() tea.Cmd {
	return nil
}

// Update implements the Bubbletea model interface
func (m selectionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i.title
				m.quitting = true
				return m, tea.Quit
			}
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		// Adjust size when terminal window changes
		m.list.SetSize(msg.Width-2, msg.Height-4)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View implements the Bubbletea model interface
func (m selectionModel) View() string {
	if m.quitting {
		return ""
	}
	return "\n" + m.list.View()
}

// InteractiveSelection presents an interactive selection prompt using Bubbletea
func InteractiveSelection(items []string, label string) (string, error) {
	// Create list items
	listItems := make([]list.Item, len(items))
	for i, s := range items {
		listItems[i] = item{title: s, desc: ""}
	}

	// Get terminal dimensions
	width, height := 80, 20 // Default reasonable size

	// Setup list configuration with delegate
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false

	// Fix spacing issue by setting the spacing to 0
	delegate.SetSpacing(0)

	// Create new list with the items and delegate
	l := list.New(listItems, delegate, width, height)

	// Configure list styles
	styles := list.DefaultStyles()
	styles.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("36")).Bold(true)
	styles.FilterCursor = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	styles.FilterPrompt = lipgloss.NewStyle().Foreground(lipgloss.Color("166"))
	styles.StatusBar = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#777777", Dark: "#999999"}).
		Padding(0, 0, 1, 2)
	styles.StatusEmpty = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	styles.NoItems = lipgloss.NewStyle().Foreground(lipgloss.Color("250")).Padding(1, 0)
	l.Styles = styles

	// Configure list options
	l.Title = label
	l.SetShowHelp(true)
	l.SetFilteringEnabled(true)

	// Initialize the model with the list
	m := selectionModel{list: l}

	// Create and run the program with alternate screen
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		return "", fmt.Errorf("error running program: %w", err)
	}

	// Return the selected item or empty string if nothing was selected
	if m.choice == "" {
		// Instead of returning an error, return the currently selected item
		// This handles cases where enter was pressed but the selected item wasn't properly captured
		if sel := l.SelectedItem(); sel != nil {
			if item, ok := sel.(item); ok {
				return item.title, nil
			}
		}
		return "", fmt.Errorf("no selection made")
	}

	return m.choice, nil
}
