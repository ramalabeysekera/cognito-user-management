package helpers

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type multiSelectModel struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	input    string
}

type indexedMultiChoice struct {
	index int
	value string
}

func initialModel(choices []string) multiSelectModel {
	return multiSelectModel{
		choices:  choices,
		selected: make(map[int]struct{}),
		input:    "",
	}
}

func (m multiSelectModel) Init() tea.Cmd {
	return nil
}

func (m multiSelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m multiSelectModel) View() string {
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

	s += "\nPress q to finish the selection.\n"
	return s
}

func (m multiSelectModel) filteredChoices() []indexedMultiChoice {
	var result []indexedMultiChoice
	for i, choice := range m.choices {
		if strings.Contains(choice, m.input) {
			result = append(result, indexedMultiChoice{i, choice})
		}
	}
	return result
}
