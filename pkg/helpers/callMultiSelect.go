package helpers

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)

func CallMultiSelect(choices []string) ([]string) {

	var result tea.Model
	var err error

	p := tea.NewProgram(initialModel(choices))
	if result, err = p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
		return nil
	}

	m := result.(multiSelectModel).selected

	var selected []string
	for i := range m {
		selected = append(selected, choices[i])
	}

	return selected
}
