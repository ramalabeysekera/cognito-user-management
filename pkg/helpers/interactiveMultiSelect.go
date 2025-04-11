package helpers

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

type item struct {
	ID         string
	IsSelected bool
}

// InteractiveMultiSelect prompts the user to select one or more items from the given list and returns the selected items.
func InteractiveMultiSelect(label string, items []string) ([]string, error) {
	const doneID = "Done"

	// Convert input items to internal item structure
	var allItems []*item
	for _, i := range items {
		allItems = append(allItems, &item{ID: i})
	}

	// Prepend "Done" item
	allItems = append([]*item{{ID: doneID}}, allItems...)

	// Recursive function for selection
	var selectItems func(selectedPos int, allItems []*item) ([]*item, error)
	selectItems = func(_ int, allItems []*item) ([]*item, error) {
		templates := &promptui.SelectTemplates{
			Label:    "{{if eq .ID \"" + doneID + "\"}}[{{ .ID | red }}]{{else}}{{if .IsSelected}}✔ {{ .ID | green }}{{else}}{{ .ID }}{{end}}{{end}}",
			Active:   "→ {{if eq .ID \"" + doneID + "\"}}[{{ .ID | red }}]{{else}}{{if .IsSelected}}✔ {{ .ID | green }}{{else}}{{ .ID }}{{end}}{{end}}",
			Inactive: "{{if eq .ID \"" + doneID + "\"}}[{{ .ID | red }}]{{else}}{{if .IsSelected}}✔ {{ .ID | green }}{{else}}{{ .ID }}{{end}}{{end}}",
		}

		prompt := promptui.Select{
			Label:     label,
			Items:     allItems,
			Templates: templates,
			Size:      10,
			Searcher: func(input string, index int) bool {
				item := strings.ToLower(allItems[index].ID)
				input = strings.ToLower(input)
				return strings.Contains(item, input)
			},
			StartInSearchMode: true,
			HideSelected:      true,
		}

		selectionIdx, _, err := prompt.Run()
		if err != nil {
			return nil, fmt.Errorf("prompt failed: %w", err)
		}

		chosenItem := allItems[selectionIdx]

		if chosenItem.ID != doneID {
			chosenItem.IsSelected = !chosenItem.IsSelected
			return selectItems(selectionIdx, allItems)
		}

		var selectedItems []*item
		for _, i := range allItems {
			if i.IsSelected {
				selectedItems = append(selectedItems, i)
			}
		}
		return selectedItems, nil
	}

	selectedItems, err := selectItems(0, allItems)
	if err != nil {
		return nil, err
	}

	// Extract IDs of selected items
	var result []string
	for _, i := range selectedItems {
		result = append(result, i.ID)
	}
	return result, nil
}
