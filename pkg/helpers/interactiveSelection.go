// Package helpers provides utility functions for interactive CLI operations
package helpers

import (
	"strings"

	"github.com/manifoldco/promptui"
)

// InteractiveSelection presents an interactive selection prompt to the user
// Parameters:
//   - list: slice of strings containing the options to choose from
//   - label: text label shown above the selection prompt
// Returns:
//   - string: the selected option
//   - error: any error that occurred during selection
func InteractiveSelection(list []string, label string) (string, error) {
	// Configure the visual templates for the prompt UI
	templates := &promptui.SelectTemplates{
		Label:    "{{ . | bold | cyan }}", // Cyan bold text for the label
		Active:   "\U0001F336 {{ . | cyan | bold }}", // Hot pepper emoji + cyan bold for active selection
		Inactive: "  {{ . | white }}", // White text for inactive items
		Selected: "\U0001F525 You chose: {{ . | green }}", // Fire emoji + green text for final selection
	}

	// Configure the selection prompt
	prompt := promptui.Select{
		Label:     label,
		Items:     list,
		Templates: templates,
		Size:      10, // Number of items to display at once
		// Define search function for filtering items
		Searcher: func(input string, index int) bool {
			item := strings.ToLower(list[index])
			input = strings.ToLower(input)
			return strings.Contains(item, input)
		},
		StartInSearchMode: true, // Start with search input active
	}

	// Run the prompt and get user selection
	_, selection, err := prompt.Run()

	// Return empty string if error occurs
	if err != nil {
		return "", nil
	}

	return selection, nil
}
