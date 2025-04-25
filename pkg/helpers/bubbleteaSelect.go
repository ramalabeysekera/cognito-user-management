package helpers

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type singleSelectModel struct {
	choices  []string         // all items
	cursor   int              // which item the cursor is pointing at (in the filtered list)
	selected map[int]struct{} // selected items (by index in the original choices)
	input    string           // user input (search)
}

// Struct to keep track of filtered item and its original index
type indexedChoice struct {
	index int
	value string
}

func initialSingleSelectModel(choices []string) singleSelectModel {
	return singleSelectModel{
		choices:  choices,
		selected: make(map[int]struct{}),
		input:    "",
	}
}

func (m singleSelectModel) Init() tea.Cmd {
	return nil
}

func (m singleSelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.filteredChoices()) - 1
			}

		case "down":
			if m.cursor < len(m.filteredChoices())-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}

		case "enter", " ":
			filtered := m.filteredChoices()
			if len(filtered) > 0 {
				actualIndex := filtered[m.cursor].index
				if _, ok := m.selected[actualIndex]; ok {
					delete(m.selected, actualIndex)
				} else {
					m.selected[actualIndex] = struct{}{}
					return m, tea.Quit
				}
			}

		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
			m.cursor = 0

		default:
			m.input += msg.String()
			m.cursor = 0
		}
	}
	return m, nil
}

func (m singleSelectModel) View() string {
	s := fmt.Sprintf("Search: %s\n\n", m.input)

	filtered := m.filteredChoices()
	for i, item := range filtered {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[item.index]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, item.value)
	}

	s += "\nPress q to quit.\n"
	return s
}

// Helper to get filtered choices with original indexes
func (m singleSelectModel) filteredChoices() []indexedChoice {
	var result []indexedChoice
	for i, choice := range m.choices {
		if strings.Contains(choice, m.input) {
			result = append(result, indexedChoice{i, choice})
		}
	}
	return result
}
