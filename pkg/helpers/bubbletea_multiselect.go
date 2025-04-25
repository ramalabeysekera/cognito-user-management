package helpers

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// multiItem represents a selectable item with multi-selection capability
type multiItem struct {
	title     string
	selected  bool
	isControl bool // For items like "Done"
}

// Title implements the list.Item interface
func (i multiItem) Title() string {
	prefix := "  "
	if i.selected {
		prefix = "✓ "
	}

	if i.isControl {
		return "[" + i.title + "]"
	}

	return prefix + i.title
}

// Description implements the list.Item interface
func (i multiItem) Description() string { return "" }

// FilterValue implements the list.Item interface
func (i multiItem) FilterValue() string { return i.title }

// multiSelectionModel represents the Bubbletea model for multiple item selection
type multiSelectionModel struct {
	list     list.Model
	choices  []string
	quitting bool
}

// Init implements the Bubbletea model interface
func (m multiSelectionModel) Init() tea.Cmd {
	return nil
}

// Update implements the Bubbletea model interface
func (m multiSelectionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			i, ok := m.list.SelectedItem().(multiItem)
			if ok && i.isControl && i.title == "Done" {
				m.choices = m.getSelectedItems()
				m.quitting = true
				return m, tea.Quit
			}
		case " ": // Space key
			// Toggle selection
			index := m.list.Index()
			items := m.list.Items()
			if index < len(items) {
				i, ok := items[index].(multiItem)
				if ok && !i.isControl {
					i.selected = !i.selected
					items[index] = i
					m.list.SetItems(items)
				}
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
func (m multiSelectionModel) View() string {
	if m.quitting {
		return ""
	}
	return "\n" + m.list.View()
}

// getSelectedItems extracts selected items from the model
func (m multiSelectionModel) getSelectedItems() []string {
	var selected []string

	for _, item := range m.list.Items() {
		i, ok := item.(multiItem)
		if ok && i.selected && !i.isControl {
			selected = append(selected, i.title)
		}
	}

	return selected
}

// InteractiveMultiSelect prompts the user to select multiple items using Bubbletea
func InteractiveMultiSelect(label string, items []string) ([]string, error) {
	// Create list items with Done control item
	listItems := make([]list.Item, len(items)+1)

	// Add "Done" control item
	listItems[0] = multiItem{title: "Done", isControl: true}

	// Add regular items
	for i, s := range items {
		listItems[i+1] = multiItem{title: s, selected: false}
	}

	// Default reasonable size
	width, height := 80, 20

	// Setup list configuration
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false

	// Fix spacing issue by setting the spacing to 0
	delegate.SetSpacing(0)

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
	m := multiSelectionModel{list: l}

	// Create and run the program with alternate screen
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		return nil, fmt.Errorf("error running program: %w", err)
	}

	// Return the selected items
	return m.choices, nil
}
