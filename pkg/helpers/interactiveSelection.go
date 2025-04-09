package helpers

import (
	"strings"

	"github.com/manifoldco/promptui"
)

func InteractiveSelection(list []string, label string) (string, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . | bold | cyan }}",
		Active:   "\U0001F336 {{ . | cyan | bold }}", // 🌶
		Inactive: "  {{ . | white }}",
		Selected: "\U0001F525 You chose: {{ . | green }}",
	}

	prompt := promptui.Select{
		Label:     label,
		Items:     list,
		Templates: templates,
		Searcher: func(input string, index int) bool {
			item := strings.ToLower(list[index])
			input = strings.ToLower(input)
			return strings.Contains(item, input)
		},
		StartInSearchMode: true,
	}

	_, selection, err := prompt.Run()

	if err != nil {
		return "", nil
	}

	return selection, nil
}
